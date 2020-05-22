package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Fatalln(http.ListenAndServe(os.Getenv("API_ADDR"), http.FileServer(http.Dir(os.Args[1]))))
}
