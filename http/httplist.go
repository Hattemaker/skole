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

func main () {
	httpAddr1 := flag.String("http", "foo", "HTTP-address") //flag for å velge http-adresse i
	// terminalen.
	flag.Parse()

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(*httpAddr1, nil)
	if err != nil {
		panic(err)
	}
	}

