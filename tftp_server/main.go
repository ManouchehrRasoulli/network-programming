package main

import (
	"flag"
	"io/ioutil"
	"log"
	"network-programming/tftp"
)

var (
	address = flag.String("a", "127.0.0.1:69", "listen address")
	payload = flag.String("p", "test.txt", "file to serve")
)

func main() {
	flag.Parse()
	p, err := ioutil.ReadFile(*payload)
	if err != nil {
		log.Panic(err)
	}

	s := tftp.Server{Payload: p}
	log.Fatal(s.ListenAndServe(*address))
}
