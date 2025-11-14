package api

import (
	"log/slog"
	"net/http"

	"go-chord-progression/audio"
	"go-chord-progression/chords"
)

const (
	webRoot        = "/"
	webResourceDir = "./web-resources"
)

// ServeAudio serves the chord progression audio file stream.
func ServeAudio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/vnd.wave")
	//w.Header().Set("Connection", "Keep-Alive")
	//w.Header().Set("Transfer-Encoding", "chunked")
	slog.Info("new connection")
	chord := chords.NewFor(chords.MajorIntervals)
	chord.Append(chords.NewFor(chords.MajorIntervals))
	chord.Append(chords.NewFor(chords.MinorMelodicIntervals))
	chord.Append(chords.NewFor(chords.MajorIntervals))
	err := audio.Generate(chord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.ServeFile(w, r, audio.OutputFile)
}

// ServeStatic serves all static files on the root route, f. i. the main page.
func ServeStatic(writer http.ResponseWriter, req *http.Request) {
	http.FileServer(http.Dir(webResourceDir)).ServeHTTP(writer, req)
}
