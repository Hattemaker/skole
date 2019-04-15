package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter port:")
	text, _ :=reader.ReadString('\n')
	fmt.Println(text)

	resp, err := http.Get("http://localhost:5000") //"http://localhost:5000"
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}