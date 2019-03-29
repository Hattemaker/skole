package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func handler(c net.Conn) {
	c.Write([]byte("ok"))
	c.Close()

}

func main() {

	////////////////////////
	type Person struct {
		Name  string
		Email string
	}

	nav := Person{Name: "Markus",
		Email: "m.sveggen@gmail.com",
	}

	b, err := json.Marshal(nav)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	//net.Conn.Write(b)

	//net.Conn.Close()

	//////////////////////////////

	fmt.Println("Starter server...")
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	fmt.Println("-- Server kjører --")
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)

	}

}
