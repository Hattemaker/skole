package main

import (
	"encoding/json"
	"fmt"
	"net"

)

	type Person struct {
		Name  string
		Email string
	}


func handler(c net.Conn) {
	var n = &Person{Name: "Markus",
		Email: "m.sveggen@gmail.com",
	}
	b, err := json.Marshal(n)
	if err != nil {
		fmt.Println("error:", err)
	}
	c.Write([]byte(b))
	c.Close()

}

func main() {


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