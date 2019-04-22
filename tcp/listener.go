package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
)

type Person struct {
	Name  string
	Email string
}

func handler(c net.Conn) {
	var n = &Person{Name: "Markus",
		Email: "m.sveggen@gmail.com",}

	b, err := json.Marshal(n)
	if err != nil {
		fmt.Println("error:", err)
	}
	c.Write([]byte(b)) //c = connection
	c.Close()
}

func main() {
	tcpAddr2 := flag.String("tcp", "foo", "a string")
	flag.Parse()

	l, err := net.Listen("tcp", *tcpAddr2) //":5000"*tcpAddr2, l = listen
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running.")

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
