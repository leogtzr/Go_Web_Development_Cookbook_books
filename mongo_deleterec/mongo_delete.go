package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	mongo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func deleteDocument(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("going to delete document in database for name:: ", name[0])
		collection := session.DB("mydb").C("employee")
		removeErr := collection.Remove(bson.M{"name": name[0]})
		if removeErr != nil {
			log.Print("error removing document from database :: ", removeErr)
			return
		}
		fmt.Fprintf(w, "Document with name '%s' has been deleted from database", name[0])
	} else {
		fmt.Fprintf(w, "Error occurre9d while deleting document in database for name :: ", name[0])
	}
}

func readDocuments(w http.ResponseWriter, r *http.Request) {
	log.Print("reading documents from database")
	var employees []Employee
	collection := session.DB("mydb").C("employee")
	err := collection.Find(bson.M{}).All(&employees)
	if err != nil {
		log.Print("error occurred while reading documents from database :: ", err)
		return
	}
	json.NewEncoder(w).Encode(employees)
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

func updateDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		employeeID, err := strconv.Atoi(id)
		if err != nil {
			log.Print("error converting string id to int ::", err)
			return
		}
		log.Print("going to update document in database for id:: ", id)
		collection := session.DB("mydb").C("employee")
		var changeInfo *mongo.ChangeInfo
		changeInfo, err = collection.Upsert(bson.M{"id": employeeID}, &Employee{employeeID, name[0]})
		if err != nil {
			log.Print("error occurred while updating record in database :: ", err)
			return
		}
		fmt.Fprintf(w, "Number of documents updated in database are :: %d", changeInfo.Updated)
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
	router.HandleFunc("/employee/delete", deleteDocument).Methods("DELETE")

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}

}
