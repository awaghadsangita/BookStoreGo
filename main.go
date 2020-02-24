package main

import (
	"awesomeProject1/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/getAllBooks",controllers.GetAllBooks)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
