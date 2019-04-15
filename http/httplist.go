package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name  string
	Email string
}

func main () {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
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