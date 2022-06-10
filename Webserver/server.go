package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	fileServer := http.FileServer(http.Dir("./htmltemplates"))
	http.Handle("/", fileServer)
}
func main() {
	//http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/form", form)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "There's an error in hello! 404Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Wrong Method!", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello from Evin's Webbie!")
}

func form(w http.ResponseWriter, r *http.Request) {

}
