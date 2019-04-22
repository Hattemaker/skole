package main

import (
	"encoding/json"
	"flag"
	"net"
)

type Person struct {
	Name  string
	Email string
}

// Skriver en JSON-struktur til serveren og oppretter en goroutine.
func handler(c net.Conn) {
	var n = &Person{Name: "Markus",
		Email: "m.sveggen@gmail.com",}

	b, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	c.Write([]byte(b))
	c.Close()
}

// Starter en tcp-server.
func main() {
	tcpAddr2 := flag.String("tcp", "foo", "a string")
	flag.Parse()

	l, err := net.Listen("tcp", *tcpAddr2) // ":5000"
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
