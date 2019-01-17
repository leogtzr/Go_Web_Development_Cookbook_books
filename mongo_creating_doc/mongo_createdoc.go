package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	mongo "gopkg.in/mgo.v2"
)

var session *mongo.Session
var connError error

// Employee ...
type Employee struct {
	ID   int    `json:"uid"`
	Name string `json:"name"`
}

func init() {
	Host := []string{
		"127.0.0.1:27017",
		// replica set addrs...
	}

	session, connError = mongo.DialWithInfo(&mongo.DialInfo{
		Addrs: Host,
		// Username: Username,
		// Password: Password,
		// Database: Database,
		// DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
		// 	return tls.Dial("tcp", addr.String(), &tls.Config{})
		// },
	})
	if connError != nil {
		panic(connError)
	}
}

func createDocument(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, nameOK := vals["name"]
	id, idOK := vals["id"]
	if nameOK && idOK {
		employeeID, err := strconv.Atoi(id[0])
		if err != nil {
			log.Print("error converting string id to int:: ", err)
			return
		}

		log.Print("going to insert document in database for name :: ", name[0])
		collection := session.DB("mydb").C("employee")
		err = collection.Insert(&Employee{employeeID, name[0]})
		if err != nil {
			log.Print("error occurred while inserting document in database :: ", err)
			return
		}
		fmt.Fprintf(w, "Last created document id is :: %s", id[0])
	} else {
		fmt.Fprint(w, "Error occurred while creating document in database for name:: ", name[0])
	}
}

func getDbNames(w http.ResponseWriter, r *http.Request) {
	db, err := session.DatabaseNames()
	if err != nil {
		log.Print("error getting database names :: ", err)
		return
	}
	fmt.Fprintf(w, "Databases names are :: %s", strings.Join(db, ", "))
}

func main() {

	defer session.Close()
	router := mux.NewRouter()
	router.HandleFunc("/employee/create", createDocument).Methods("POST")

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}

}
