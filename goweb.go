package main

import (
	"fmt"
	"log"
	"net/http"
) //library for networking

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hii, saurabhs051 !!")
}


func main() {
	http.HandleFunc("/", index_handler) 
	err := http.ListenAndServe(":8080", nil) 
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}

}
