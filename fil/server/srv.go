package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"network-programming/fil/proto"
)

type Server struct {
	listener net.Listener
	Address  string
	Logger   *log.Logger
	Exit     chan struct{}
}

func (s *Server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}

	s.listener = lis

	go func(ctx context.Context) { // run program in background
		s.Logger.Print("server :: ready for accepting connections...\n")
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				s.Logger.Printf("server :: got error \"%v\" on listening\n", err)
			}

			go func(conn net.Conn) {
				s.serve(conn) // handle connection
				s.Logger.Printf("server :: connection closed %s\n", conn.RemoteAddr())
			}(conn)
		} // in case of any error
		s.Exit <- struct{}{}
	}(ctx)

	s.Logger.Printf("server :: listening on address :: %s\n", lis.Addr())
	return nil
}

func (s *Server) Wait() {
	<-s.Exit
}

// serve
// will accept connection and will manage data transfer within given connection.
func (s *Server) serve(con net.Conn) {
	s.Logger.Printf("server :: accept connection %v\n", con.RemoteAddr())

	reader := bufio.NewReader(con)
	req, err := ioutil.ReadAll(reader)
	if err != nil {
		s.Logger.Printf("server :: got error %s on connection %v\n", err, con.RemoteAddr())
	}

	d := proto.RReq{}
	if err := d.UnmarshalBinary(req); err != nil {
		s.Logger.Printf("server :: got error %s on connection %v\n", err, con.RemoteAddr())
	}

	fmt.Printf("%v\n", d)
}
