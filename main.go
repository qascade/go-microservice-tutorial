package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})

	http.HandleFunc("/percole", func(http.ResponseWriter, *http.Request) {
		log.Println("percole")
	})


	http.ListenAndServe(":9090", nil)
}