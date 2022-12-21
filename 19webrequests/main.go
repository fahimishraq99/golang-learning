package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://bracu.ac.bd"

func main() {
	fmt.Println("Welcome to class on web request")

	response, err := http.Get(url) //making http request

	if err != nil { //error checking
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body) //reading the response all by 'ioutil'

	if err != nil {
		panic(err)
	}
	content := string(databytes) //conversion of content
	fmt.Println(content)
}
