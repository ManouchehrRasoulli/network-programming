package server

import "net"

type Server struct {
	listener net.Listener
	address  string
}
