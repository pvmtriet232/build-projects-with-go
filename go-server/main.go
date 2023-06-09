package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("hello", helloHandler)
	fmt.Printf("startin server at port 8080\n")
	if err := http.ListenAndServe("8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloHandler")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello formHandler")
}
