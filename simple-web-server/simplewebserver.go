package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_PORT = "8080"
	CONN_HOST = "localhost"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
