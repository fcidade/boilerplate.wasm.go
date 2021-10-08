package main

import (
	"log"
	"net/http"
)

const (
	addr       = ":9000"
	assetsPath = "./assets"
)

func main() {
	if err := http.ListenAndServe(addr, http.FileServer(http.Dir(assetsPath))); err != nil {
		log.Fatalln("Failed to start server", err)
	}
}
