package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"

	"go-chord-progression/audio"
	"go-chord-progression/chords"

	"github.com/0x6flab/namegenerator"
)

const (
	webRoot        = "/"
	webResourceDir = "./web-resources"
)

const outputFilePath = "/home/bschaa/projekte/privat/2025-go-chord-progression/target/"

func GenerateNewFileEndpoint(w http.ResponseWriter, r *http.Request) {
	outputFile, err := createAudioFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultHtml := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<header>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta title="go-chord-progression"/>
    <link rel="stylesheet" href="_reset.css">
    <link rel="stylesheet" href="player.css"/>
</header>
<body>

<h1>Go Chord Progression</h1>
<p>
<form action="http://localhost:9001/generate" method="post" >
    <button id="mybutton" type="submit" name="createNewSong">Create a new song</button>
</form>
</p>

	<audio id="player" controls autoplay>
        <!-- workaround: if you don't use the default port 9001, you must set the actual port here too -->
        <source src="http://localhost:9001/audio?file=%s" type="audio/wav">
        <h1>Oh no, something wrong with your browser?</h1>
    	<p>It looks like that your browser does not support HTML5 audio players. In 2025 and later? Weird.</p>
	</audio>
<script type="text/javascript">
	document.getElementById("player").volume=0.2;
</script>
</body>
</html>
`, outputFile)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write([]byte(resultHtml))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ServeAudio serves the chord progression audio file stream.
func ServeAudio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/vnd.wave")
	//w.Header().Set("Connection", "Keep-Alive")
	//w.Header().Set("Transfer-Encoding", "chunked")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outputFile := r.FormValue("file")
	slog.Info("new connection for " + outputFile)

	http.ServeFile(w, r, filepath.Join(outputFilePath, outputFile))
}

// ServeStatic serves all static files on the root route, f. i. the main page.
func ServeStatic(writer http.ResponseWriter, req *http.Request) {
	http.FileServer(http.Dir(webResourceDir)).ServeHTTP(writer, req)
}

func createAudioFile() (string, error) {
	filename := namegenerator.NewGenerator().Generate()

	chord := chords.NewFor(chords.MajorIntervals)
	chord.Append(chords.NewFor(chords.MajorIntervals))
	chord.Append(chords.NewFor(chords.MinorMelodicIntervals))
	chord.Append(chords.NewFor(chords.MajorIntervals))

	outputFile := filepath.Join(outputFilePath, filename+".wav")
	err := audio.Generate(chord, outputFile)
	if err != nil {
		return "", fmt.Errorf("failed to generate file %s: %w", filename, err)
	}
	return filename + ".wav", nil
}
