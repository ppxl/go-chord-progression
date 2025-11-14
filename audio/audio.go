package audio

import (
	"fmt"

	"go-chord-progression/chords"

	synth "github.com/DylanMeeus/GoAudio/synthesizer"
	"github.com/DylanMeeus/GoAudio/wave"
)

const (
	sine = iota
	square
	downsaw
	upsaw
	triangle
)

const (
	sampleRate     = 44100
	durationInSecs = sampleRate * 1
)

// Generate creates an audio wave from an oscillator of a given shape and durationInSecs
// Modified by amplitude / frequency breakpoints
func Generate(theChords *chords.Chord, filename string) error {

	osc, err := synth.NewOscillator(sampleRate, synth.SQUARE)
	if err != nil {
		return fmt.Errorf("failed to create oscillator: %w", err)
	}

	frames := []wave.Frame{}
	var timeframe int
	for count := 0; count < theChords.Count(); count++ {
		for durationTick := 0; durationTick <= durationInSecs; durationTick++ {
			a := float64(1)
			d := float64(0.5)
			s := float64(0.3)
			r := float64(0.1)
			const maxamp = 0.6
			value := synth.ADSR(maxamp, durationInSecs, a, d, s, r, sampleRate, timeframe)

			chordBaseFreq := theChords.Notes[count].BaseFrequency
			frequencyStep := value * osc.Tick(chordBaseFreq)
			frames = append(frames, wave.Frame(frequencyStep))
			timeframe++
		}
	}

	bitsPerSample := 16
	wfmt := wave.NewWaveFmt(1, 2, sampleRate, bitsPerSample, nil)
	err = wave.WriteFrames(frames, wfmt, filename)
	if err != nil {
		return fmt.Errorf("failed to write wave to file %s: %w", filename, err)
	}

	return nil
}

// FramesToFloats turns a slice of frames into a slice of float64.
func framesToFloats(frames []wave.Frame) []float64 {
	floats := make([]float64, len(frames))
	for i, frame := range frames {
		floats[i] = float64(frame)
	}
	return floats
}
