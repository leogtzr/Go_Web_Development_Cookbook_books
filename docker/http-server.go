package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	ConnHost = "0.0.0.0"
	ConnPort = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(ConnHost+":"+ConnPort, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	resp, err := http.Get("https://google.com")
// 	check(err)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	check(err)
// 	fmt.Println(string(body))
// }

// func check(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }
