package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Person1 struct {
	Name  string
	Email string
}

var person1 Person1

func writeToFile () {
	info := "\nNavn: " + person1.Name + " - Epost: " + person1.Email //variabel med oppsett til fil

	file, err := os.OpenFile("http.txt", os.O_WRONLY|os.O_APPEND, 0644)// Åpner filen som write-only
	// og tilføyer filen data når man skriver til filen "http.txt".
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(info)) //skriver slicen med bytes til filen ("http.txt")
	if err != nil {
		log.Fatal(err)
	}

}

func main () 		{
	httpAddr2  := flag.String("http", "foo", "HTTP address")
	flag.Parse()

	resp, err := http.Get(*httpAddr2) //"http://localhost:5000"
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &person1); err != nil { //bs er data lest fra servern i bytes, person1 er struct oppsett
		panic(err)
	}
	writeToFile()
}
