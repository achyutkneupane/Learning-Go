// basic http server in go to write Hello World

package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("localhost:5001", nil))

}