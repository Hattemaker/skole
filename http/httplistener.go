package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Person struct {
	Name  string
	Email string
}

// Skriver en JSON-struktur til serveren.
func handler (w http.ResponseWriter, r *http.Request) {
	var n1 = &Person{Name: "Markus",
		Email: "m.sveggen@gmail.com",}

	js, err := json.Marshal(n1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("JSON-struct received")
	w.Write(js)

}

// Starter en HTTP-server.
func main () {
	httpAddr1 := flag.String("http", "foo", "HTTP-address")
	flag.Parse()

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(*httpAddr1, nil) // "http://localhost:5000"
	if err != nil {
		panic(err)
	}
}

