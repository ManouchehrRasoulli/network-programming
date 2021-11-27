package tftp

import (
	"bytes"
	"errors"
	"log"
	"net"
	"time"
)

type Server struct {
	Payload []byte        // at this point we only return thus byte array as server response
	Retries uint8         // number of server retries
	Timeout time.Duration // ack timeout
}

func (s Server) ListenAndServe(add string) error {
	conn, err := net.ListenPacket("udp", add)
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()

	log.Printf("listening on port %d ...", conn.LocalAddr())
	return s.Serve(conn)
}

func (s *Server) Serve(conn net.PacketConn) error {
	if conn == nil {
		return errors.New("nil connection")
	}

	if s.Payload == nil {
		return errors.New("payload is required")
	}

	if s.Retries == 0 {
		s.Retries = 10
	}

	if s.Timeout == 0 {
		s.Timeout = time.Second * 6 // 6sec for timeout
	}

	var rrq ReadRequest
	for {
		buff := make([]byte, DatagramSize)
		_, add, err := conn.ReadFrom(buff)
		if err != nil {
			return err
		}

		err = rrq.UnmarshalBinary(buff)
		if err != nil {
			log.Printf("[%s] bad request: %v", add, err)
			continue
		}

		go s.handle(add.String(), rrq) // give a copy to handler to manage received request
	}
}

func (s *Server) handle(addr string, rrq ReadRequest) {
	log.Printf("[%s] requested file : %s", addr, rrq.FileName)

	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Printf("[%s] dial: %v", addr, err)
		return
	}

	var (
		ackPkt  Ack
		errPkt  Err
		dataPkt = Data{Payload: bytes.NewReader(s.Payload)}
		buffer  = make([]byte, DatagramSize)
	)

	{ // send process
	NEXTPACKET:
		for n := DatagramSize; n == DatagramSize; {
			data, err := dataPkt.MarshalBinary()
			if err != nil {
				log.Printf("[%s] preparing data packet : %v", addr, err)
				return
			}

			{ // retry
				for i := s.Retries; i > 0; i-- {
					n, err = conn.Write(data)
					if err != nil {
						log.Printf("[%s] write: %v", addr, err)
						return
					}

					_ = conn.SetReadDeadline(time.Now().Add(s.Timeout))

					_, err := conn.Read(buffer)
					if err != nil {
						if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
							continue // return to retry
						}

						log.Printf("[%s] waiting for ACK: %v", addr, err)
						return
					}

					switch {
					case ackPkt.UnmarshalBinary(buffer) == nil:
						if uint16(ackPkt) == dataPkt.BlockNumber {
							continue NEXTPACKET
						}
					case errPkt.UnmarshalBinary(buffer) == nil:
						log.Printf("[%s] receive error: %v", addr, errPkt)
						return
					default:
						log.Printf("[%s] bad packet", addr)
					}
				}

				log.Printf("[%s] exhausted retrues", addr)
				return
			}
		}

		log.Printf("[%s] sent %d blocks", addr, dataPkt.BlockNumber)
	}
}
