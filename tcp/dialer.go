package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type Person1 struct {
	Name  string
	Email string
}

var person1 Person1

func WriteToFile () {
	info := "\nNavn: " + person1.Name + " - Epost: " + person1.Email

	file, err := os.OpenFile("tcp.txt", os.O_WRONLY|os.O_APPEND, 0644) // Åpner filen som write-only
	// og tilføyer filen data når man skriver til filen "tcp.txt".
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lenk, err := file.Write([]byte(info)) //skriver slicen med bytes til filen ("http.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lenk)
}

func main () {
	tcpAddr1 := flag.String("tcp", "foo", "a string")
	flag.Parse()

	conn, err := net.Dial("tcp", *tcpAddr1) //":5000"
	if err != nil {
		panic(err)
	}
	fmt.Println("Dialer running.")
	defer conn.Close() //Utsetter å lukke tilkoblingen med serveren.

	bs, _ := ioutil.ReadAll(conn) //Leser
	if err := json.Unmarshal(bs, &person1); err != nil { //bs er data lest fra servern i bytes, person1 er struct oppsett
		panic(err)
	}
	fmt.Println(person1.Name, "-", person1.Email) //Printer json structuren i unmarshaled format.
	WriteToFile()
}
