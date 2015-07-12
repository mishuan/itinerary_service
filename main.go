package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Server started and listening to port 8080.\n")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
