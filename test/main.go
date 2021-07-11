package main

import (
	"log"
	"net/http"
	"osqueryi/req"
)

func main() {
	http.HandleFunc("/", req.Handler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Println(err)
	}
}
