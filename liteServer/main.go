package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
)

func main() {
	workPath := "./"
	if len(os.Args) > 1 {
		fmt.Println("Usage: liteServer <Server Location>")
		workPath = os.Args[1]
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(workPath))))

	server := &http.Server{
		Addr:         "localhost:1080",
		Handler:      mux,
	}
	log.Println("liteServer is starting >>>")
	log.Printf("Running at %v\n", server.Addr)
	log.Printf("Working at %s\n", workPath)
	server.ListenAndServe()
}