package api

import (
	"fmt"
	"net/http"
	"time"
)

const (
	webRoot        = "/"
	webResourceDir = "./web-resources"
)

// ServeAudio serves the chord progression audio file stream.
func ServeAudio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/vnd.wave")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")

	for i := 0; i < 30; i++ {
		_, err := fmt.Fprintf(w, "data: %d\n", i) // stream audio here
		if err != nil {
			http.Error(w, "error while writing to streaming", http.StatusInternalServerError)
		}
		flusher.Flush()
		time.Sleep(300 * time.Millisecond)
	}
	// fixme, read file into a buffer and try to stream it slowly
	http.ServeFile(w, r, "sample.mp3")
}

// ServeStatic serves all static files on the root route, f. i. the main page.
func ServeStatic(writer http.ResponseWriter, req *http.Request) {
	http.FileServer(http.Dir(webResourceDir)).ServeHTTP(writer, req)
}
