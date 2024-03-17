package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Register the route
	http.HandleFunc("/", helloHandler)

	// Start the server
	http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func getPort() string {
	if v, ok := os.LookupEnv("APP_PORT"); ok {
		return v
	}
	return "80"
}
