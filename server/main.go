package main

import (
	"log"
	"net/http"
)

func main() {
	fileSrv := http.FileServer(http.Dir("./"))

	err := http.ListenAndServe(":8083", fileSrv)
	if err != nil {
		log.Fatalln(err)
	}
}
