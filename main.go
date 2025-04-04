package main

import (
	"fmt"
	"net/http"
	"os"
	"untitled/api"
)

const defaultPort = "9001"

func main() {
	http.HandleFunc("/", api.ServeStatic)
	http.HandleFunc("/audio", api.ServeAudio)

	port, isSet := os.LookupEnv("CHORD_PROG_PORT")
	if isSet {
		fmt.Println("Non-default port found. If not already done, you may want to update the port in the index.html file")
	} else {
		port = defaultPort
	}
	fmt.Printf("Listening on port %s\n", port)

	fmt.Printf("Point your browser to http://localhost:%s/index.html\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(fmt.Errorf("failed to serve on port %s: %w", port, err))
	}
}
