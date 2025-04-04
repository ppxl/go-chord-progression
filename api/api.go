package api

import (
	"net/http"
)

const (
	webRoot        = "/"
	webResourceDir = "./web-resources"
)

// ServeAudio serves the chord progression audio file stream.
func ServeAudio(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "sample.wav")
}

// ServeStatic serves all static files on the root route, f. i. the main page.
func ServeStatic(writer http.ResponseWriter, req *http.Request) {
	http.FileServer(http.Dir(webResourceDir)).ServeHTTP(writer, req)
}
