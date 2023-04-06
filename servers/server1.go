package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Port
	PORT := "5051"
	//Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user", userHandler)

	//Start server
	fmt.Println("Initial server on port: ", PORT)
	http.ListenAndServe(":5051", nil)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, world!"))
}

func userHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome user!"))
}
