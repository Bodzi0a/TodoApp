package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Working"))
}

func main() {
	http.HandleFunc("/health", healthHandler)
	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
