package main

import "log"

func main() {
	c := Client{
		Address: "localhost:62715",
		Logger:  log.Default(),
	}

	err := c.Download("test.txt")
	if err != nil {
		panic(err)
	}
}
