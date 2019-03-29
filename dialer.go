package main

//dxkdkd

import "net"
import "io/ioutil"
import "fmt"

func main() {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close() //Utsetter å lukke tilkoblingen med serveren.

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs)) // Konverterer slicen med bytes og konverterer den til String

}
