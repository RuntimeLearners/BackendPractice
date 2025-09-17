package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error, cannot start server")
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Backend Process Success")
}
