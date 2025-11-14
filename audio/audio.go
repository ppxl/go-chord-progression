package audio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"go-chord-progression/chords"

	"github.com/DylanMeeus/GoAudio/breakpoint"
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
			const maxamp = 1
			value := synth.ADSR(maxamp, durationInSecs, a, d, s, r, sampleRate, timeframe)

			timeframe++
			//floats := FramesToFloats(frames)
			//filteredValue := synth.Lowpass(floats, 500.0, 3, sampleRate)
			//frames = wave.FloatsToFrames(filteredValue)
			frames = append(frames, wave.Frame(value*osc.Tick(theChords.Notes[count].BaseFrequency)))
		}
	}

	bitsPerSample := 16
	wfmt := wave.NewWaveFmt(1, 1, sampleRate, bitsPerSample, nil)
	err = wave.WriteFrames(frames, wfmt, filename)
	//bytes := toBytes(frames, wfmt)
	if err != nil {
		return fmt.Errorf("failed to write wave to file %s: %w", filename, err)
	}

	return nil

	fmt.Println("done")
	return nil
}

func toBytes(samples []wave.Frame, wfmt wave.WaveFmt) []byte {

	// construct this in reverse (Data -> Fmt -> Header)
	// as Fmt needs info of Data, and Hdr needs to know entire length of file

	// write chunkSize
	bits := []byte{}

	wfb := fmtToBytes(wfmt)
	data, databits := framesToData(samples, wfmt)
	hdr := createHeader(data)

	bits = append(bits, hdr...)
	bits = append(bits, wfb...)
	bits = append(bits, databits...)
	return bits
}

func createFrequencyPointStream(err error, wfmt wave.WaveFmt, tonebrks breakpoint.Breakpoints) (*breakpoint.BreakpointStream, error) {
	//fixme
	//var freqpoints string
	//freqs, err := os.ReadFile(freqpoints)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to read freq points: %w", err)
	//}
	//freqs := []byte(`
	//0:100
	//10:50
	//20:100
	//`)
	//freqPoints, err := breakpoint.ParseBreakpoints(bytes.NewReader(freqs))
	//if err != nil {
	//	return nil, fmt.Errorf("failed to parse breakpoints: %w", err)
	//}
	freqStream, err := breakpoint.NewBreakpointStream(tonebrks, wfmt.SampleRate)
	if err != nil {
		return nil, fmt.Errorf("failed to create breakpoint stream: %w", err)
	}
	return freqStream, nil
}

func createBreakpointStream(wfmt wave.WaveFmt) (*breakpoint.BreakpointStream, error) {
	////fixme
	//var amppoints string
	//amps, err := os.ReadFile(amppoints)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create wave format: %w", err)
	//}
	amps := []byte(`
	0:100
	10:100
	20:100
	`)
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

func int16ToBytes(i int) []byte {
	b := make([]byte, 2)
	in := uint16(i)
	binary.LittleEndian.PutUint16(b, in)
	return b
}

func int32ToBytes(i int) []byte {
	b := make([]byte, 4)
	in := uint32(i)
	binary.LittleEndian.PutUint32(b, in)
	return b
}

type intsToBytesFunc func(i int) []byte

var (
	// intsToBytesFm to map X-bit int to byte functions
	intsToBytesFm = map[int]intsToBytesFunc{
		16: int16ToBytes,
		32: int32ToBytes,
	}
)
var (
	// figure out which 'to int' function to use..
	byteSizeToIntFunc = map[int]bytesToIntF{
		16: bits16ToInt,
		32: bits32ToInt,
	}

	byteSizeToFloatFunc = map[int]bytesToFloatF{
		16: bitsToFloat,
		32: bitsToFloat,
		64: bitsToFloat,
	}

	// max value depending on the bit size
	maxValues = map[int]int{
		8:  math.MaxInt8,
		16: math.MaxInt16,
		32: math.MaxInt32,
		64: math.MaxInt64,
	}
)

// for our wave format we expect double precision floats
func bitsToFloat(b []byte) float64 {
	var bits uint64
	switch len(b) {
	case 2:
		bits = uint64(binary.LittleEndian.Uint16(b))
	case 4:
		bits = uint64(binary.LittleEndian.Uint32(b))
	case 8:
		bits = binary.LittleEndian.Uint64(b)
	default:
		panic("Can't parse to float..")
	}
	float := math.Float64frombits(bits)
	return float
}
func bits16ToInt(b []byte) int {
	if len(b) != 2 {
		panic("Expected size 4!")
	}
	var payload int16
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &payload)
	if err != nil {
		// TODO: make safe
		panic(err)
	}
	return int(payload) // easier to work with ints
}

// turn a 32-bit byte array into an int
func bits32ToInt(b []byte) int {
	if len(b) != 4 {
		panic("Expected size 4!")
	}
	var payload int32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &payload)
	if err != nil {
		// TODO: make safe
		panic(err)
	}
	return int(payload) // easier to work with ints
}
func framesToData(frames []wave.Frame, wfmt wave.WaveFmt) (wave.WaveData, []byte) {
	b := []byte{}
	raw := samplesToRawData(frames, wfmt)

	// We receive frames but have to store the size of the samples
	// The size of the samples is frames / channels..
	subchunksize := (len(frames) * wfmt.NumChannels * wfmt.BitsPerSample) / 8
	subBytes := int32ToBytes(subchunksize)

	// construct the data part..
	b = append(b, wave.Subchunk2ID...)
	b = append(b, subBytes...)
	b = append(b, raw...)

	wd := wave.WaveData{
		Subchunk2ID:   wave.Subchunk2ID,
		Subchunk2Size: subchunksize,
		RawData:       raw,
		Frames:        frames,
	}
	return wd, b
}

func floatToBytes(f float64, nBytes int) []byte {
	bits := math.Float64bits(f)
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, bits)
	// trim padding
	switch nBytes {
	case 2:
		return bs[:2]
	case 4:
		return bs[:4]
	}
	return bs
}

// Turn the samples into raw data...
func samplesToRawData(samples []wave.Frame, props wave.WaveFmt) []byte {
	raw := []byte{}
	for _, s := range samples {
		// the samples are scaled - rescale them?
		rescaled := rescaleFrame(s, props.BitsPerSample)
		bits := intsToBytesFm[props.BitsPerSample](rescaled)
		raw = append(raw, bits...)
	}
	return raw
}

// rescale frames back to the original values..
func rescaleFrame(s wave.Frame, bits int) int {
	rescaled := float64(s) * float64(maxValues[bits])
	return int(rescaled)
}

func fmtToBytes(wfmt wave.WaveFmt) []byte {
	b := []byte{}

	subchunksize := int32ToBytes(wfmt.Subchunk1Size)
	audioformat := int16ToBytes(wfmt.AudioFormat)
	numchans := int16ToBytes(wfmt.NumChannels)
	sr := int32ToBytes(wfmt.SampleRate)
	br := int32ToBytes(wfmt.ByteRate)
	blockalign := int16ToBytes(wfmt.BlockAlign)
	bitsPerSample := int16ToBytes(wfmt.BitsPerSample)

	b = append(b, wfmt.Subchunk1ID...)
	b = append(b, subchunksize...)
	b = append(b, audioformat...)
	b = append(b, numchans...)
	b = append(b, sr...)
	b = append(b, br...)
	b = append(b, blockalign...)
	b = append(b, bitsPerSample...)

	return b
}

// turn the sample to a valid header
func createHeader(wd wave.WaveData) []byte {
	// write chunkID
	bits := []byte{}

	chunksize := 36 + wd.Subchunk2Size
	cb := int32ToBytes(chunksize)

	bits = append(bits, wave.ChunkID...) // in theory switch on endianness..
	bits = append(bits, cb...)
	bits = append(bits, wave.ChunkID...)

	return bits
}

type (
	bytesToIntF   func([]byte) int
	bytesToFloatF func([]byte) float64
)

// FramesToFloats turns a slice of frames into a slice of float64.
func FramesToFloats(frames []wave.Frame) []float64 {
	floats := make([]float64, len(frames))
	for i, frame := range frames {
		floats[i] = float64(frame)
	}
	return floats
}
