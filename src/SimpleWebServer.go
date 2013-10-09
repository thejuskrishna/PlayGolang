package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {

	log.Println("Starting webserver")
	http.HandleFunc("/", PrintName)
	http.Handle("/output/", http.StripPrefix("/output/", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8080", nil)

}

func PrintName(w http.ResponseWriter, r *http.Request) {
	var name string
	name = r.FormValue("name")
	fmt.Fprintf(w, "Name : %s", name)
	log.Printf("Request received [%v]\n", r)
}