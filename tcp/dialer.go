package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter port:")
	text, _ :=reader.ReadString('\n')
	fmt.Println(text)

	conn, err := net.Dial("tcp", ":5000") //:5000
	if err != nil {
		panic(err)
	}
	fmt.Println("Dialer running.")
	defer conn.Close() //Utsetter å lukke tilkoblingen med serveren.

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs)) // Konverterer slicen med bytes og konverterer den til String



}

