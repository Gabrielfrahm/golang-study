package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ola mundo"))
	})
	log.Fatal(http.ListenAndServe(":3333", nil))
}
