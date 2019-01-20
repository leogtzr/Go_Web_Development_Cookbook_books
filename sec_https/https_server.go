package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	ConnHost         = "localhost"
	ConnPort         = "8443"
	HttpsCertificate = "domain.crt"
	DomainPrivateKey = "domain.key"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hllo World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServeTLS(ConnHost+":"+ConnPort, HttpsCertificate, DomainPrivateKey, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}
