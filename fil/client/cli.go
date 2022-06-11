package main

import (
	"log"
	"net"
	"network-programming/fil/proto"
)

type Client struct {
	Address string
	conn    net.Conn
	Logger  *log.Logger
}

func (c *Client) Download(path string) error {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		return err
	}

	c.Logger.Printf("client address is %s\n", conn.LocalAddr())

	req := proto.RReq{
		Path: path,
		Seq:  2,
		Size: 512,
	}

	bin, err := req.MarshalBinary()
	if err != nil {
		return err
	}

	n, err := conn.Write(bin)
	if err != nil {
		return err
	}

	if n != len(bin) {
		panic("invalid write.")
	}

	return nil
}
