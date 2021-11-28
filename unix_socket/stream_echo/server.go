package stream_echo

import (
	"context"
	"net"
)

// streamEchoServer
// the stream socket will accept type of network and corresponding address and will work
// on unix or udp connections
func streamEchoServer(ctx context.Context, network string, addr string) (net.Addr, error) {
	server, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = server.Close()
		}()

		for {
			conn, err := server.Accept()
			if err != nil {
				return
			}

			go func() { // manage accepted connection
				defer func() {
					_ = conn.Close()
				}()

				for {
					buff := make([]byte, 1024)
					n, err := conn.Read(buff)
					if err != nil {
						return
					}

					_, err = conn.Write(buff[:n])
					if err != nil {
						return
					}
				}
			}()
		}
	}()

	return server.Addr(), nil
}
