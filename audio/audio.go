package audio

import (
	"bytes"
	"fmt"
	"os"

	"github.com/DylanMeeus/GoAudio/breakpoint"
	synth "github.com/DylanMeeus/GoAudio/synthesizer"
	"github.com/DylanMeeus/GoAudio/wave"
)

var stringToShape = map[string]synth.Shape{
	"sine":     0,
	"square":   1,
	"downsaw":  2,
	"upsaw":    3,
	"triangle": 4,
}

// example use of the oscillator to generate different waveforms
var (
	duration = 10
	shape    = "triangle"
	output   = "audio.mp3"
)

// Generate creates an audio wave from an oscillator of a given shape and duration
// Modified by amplitude / frequency breakpoints
func Generate() error {
	// create wave file sampled at 44.1Khz w/ 16-bit frames
	wfmt := wave.NewWaveFmt(1, 1, 44100, 16, nil)
	ampStream, err := createBreakpointStream(wfmt)
	if err != nil {
		return err
	}

	freqStream, err := createFrequencyPointStream(err, wfmt)
	if err != nil {
		return err
	}

	frames, err := generate(duration, stringToShape[shape], ampStream, freqStream, wfmt)
	if err != nil {
		return fmt.Errorf("failed to generate audio frames: %w", err)
	}

	err = wave.WriteFrames(frames, wfmt, output)
	if err != nil {
		return fmt.Errorf("failed to write audio frames: %w", err)
	}

	fmt.Println("done")
	return nil
}

func createFrequencyPointStream(err error, wfmt wave.WaveFmt) (*breakpoint.BreakpointStream, error) {
	//fixme
	var freqpoints string
	freqs, err := os.ReadFile(freqpoints)
	if err != nil {
		return nil, fmt.Errorf("failed to read freq points: %w", err)
	}
	freqPoints, err := breakpoint.ParseBreakpoints(bytes.NewReader(freqs))
	if err != nil {
		return nil, fmt.Errorf("failed to parse breakpoints: %w", err)
	}
	freqStream, err := breakpoint.NewBreakpointStream(freqPoints, wfmt.SampleRate)
	if err != nil {
		return nil, fmt.Errorf("failed to create breakpoint stream: %w", err)
	}
	return freqStream, nil
}

func createBreakpointStream(wfmt wave.WaveFmt) (*breakpoint.BreakpointStream, error) {
	//fixme
	var amppoints string
	amps, err := os.ReadFile(amppoints)
	if err != nil {
		return nil, fmt.Errorf("failed to create wave format: %w", err)
	}
	ampPoints, err := breakpoint.ParseBreakpoints(bytes.NewReader(amps))
	if err != nil {
		return nil, fmt.Errorf("failed to parse breakpoints: %w", err)
	}
	ampStream, err := breakpoint.NewBreakpointStream(ampPoints, wfmt.SampleRate)
	if err != nil {
		return nil, fmt.Errorf("failed to create breakpoint stream: %w", err)
	}
	return ampStream, nil
}

func generate(dur int, shape synth.Shape, ampStream, freqStream *breakpoint.BreakpointStream, wfmt wave.WaveFmt) ([]wave.Frame, error) {
	reqFrames := dur * wfmt.SampleRate
	frames := make([]wave.Frame, reqFrames)
	osc, err := synth.NewOscillator(wfmt.SampleRate, shape)
	if err != nil {
		return nil, fmt.Errorf("failed to create oscillator: %w", err)
	}

	for i := range frames {
		amp := ampStream.Tick()
		freq := freqStream.Tick()
		frames[i] = wave.Frame(amp * osc.Tick(freq))
	}

	return frames, nil
}
