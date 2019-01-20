package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	ConnHost           = "localhost"
	ConnPort           = "8080"
	ClaimIssuer        = "Packt"
	ClaimExpiryInHours = 24
)

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees Employees
var signature = []byte("secret")

var jwtMiddleWare = jwtmiddleware.New(
	jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return signature, nil
		}, SigningMethod: jwt.SigningMethodHS256,
	},
)

func init() {
	employees = Employees{
		Employee{ID: 1, FirstName: "Foo", LastName: "Bar"},
		Employee{ID: 2, FirstName: "Baz", LastName: "Qux"},
	}
}

func getToken(w http.ResponseWriter, r *http.Request) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * ClaimExpiryInHours).Unix(),
		Issuer:    ClaimIssuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(signature)
	w.Write([]byte(tokenString))
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/status", getStatus).Methods("GET")
	muxRouter.HandleFunc("/get-token", getToken).Methods("GET")
	// muxRouter.HandleFunc("/employees", getEmployees).Methods("GET")
	muxRouter.Handle("/employees", jwtMiddleWare.Handler(http.HandlerFunc(getEmployees))).Methods("GET")

	err := http.ListenAndServe(ConnHost+":"+ConnPort, handlers.LoggingHandler(os.Stdout, muxRouter))
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}

}
