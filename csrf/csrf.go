package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

const (
	ConnHost         = "localhost"
	ConnPort         = "8443"
	HttpsCertificate = "domain.crt"
	DomainPrivateKey = "domain.key"
)

var authKey = []byte("authentication-key")

func signUp(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("sign-up.html")
	err := parsedTemplate.Execute(w, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
	if err != nil {
		log.Print("Error occurred while executing the template :: ", err)
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print("error occurred while parsing from ", err)
	}
	name := r.FormValue("name")
	fmt.Fprintf(w, "Hi %s", name)
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/signup", signUp)
	muxRouter.HandleFunc("/post", post)

	protect := csrf.Protect(authKey)

	err := http.ListenAndServeTLS(ConnHost+":"+ConnPort, HttpsCertificate, DomainPrivateKey, protect(muxRouter))
	if err != nil {
		log.Fatal("error starting http server :: ", err)
	}
}
