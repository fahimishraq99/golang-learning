package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fahimishraq99/mongoapi/router"
)

func main() {
	fmt.Println("Welcome to class on mongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
