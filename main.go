package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mouni0415/testing-http-server/handlers"
)

func main() {
	//define routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/users", handlers.UserHandler)

	port := "8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
