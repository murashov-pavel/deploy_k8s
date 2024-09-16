package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is a testing app!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /health request\n")
	w.WriteHeader(200)
	io.WriteString(w, "I am alive!\n")
}

func getSecure(w http.ResponseWriter, r *http.Request) {
	secureVar, exists := os.LookupEnv("SECURE_VAR")
	fmt.Printf("got /secure request, ")
	if exists && secureVar == "123" {
		fmt.Printf("Secure var found\n")
		w.WriteHeader(200)
		io.WriteString(w, "Access granted\n")

	} else {
		fmt.Printf("Secure var not found\n")
		w.WriteHeader(401)
		io.WriteString(w, "Access denied\n")

	}
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /config request\n")
	config, err := os.ReadFile("./config.yaml")
	if err != nil {
		w.WriteHeader(404)
		io.WriteString(w, "404, Config file not found")
	} else {
		io.WriteString(w, string(config))
	}

}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/health", getHello)
	http.HandleFunc("/secure", getSecure)
	http.HandleFunc("/config", getConfig)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server\n")
		os.Exit(1)
	}
}
