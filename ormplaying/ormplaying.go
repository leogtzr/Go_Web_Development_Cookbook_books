package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	mongo "gopkg.in/mgo.v2"
)

var session *mongo.Session
var connError error

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

	http.HandleFunc("/", getDbNames)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}

}
