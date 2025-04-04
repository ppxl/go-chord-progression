package chords

import (
	"fmt"
)

// note frequencies in Hertz (Hz), based on DIN Pitch Standard tone A, 440 Hz
var (
	NoteA__0 = Note{"A__0", 440.00}
	NoteAis0 = Note{"Ais0", 466.16}
	NoteB__0 = Note{"B__0", 493.88}
	NoteC__0 = Note{"C__0", 523.25}
	NoteCis0 = Note{"Cis0", 554.37}
	NoteD__0 = Note{"D__0", 587.33}
	NoteDis0 = Note{"Dis0", 622.25}
	NoteE__0 = Note{"E__0", 659.26}
	NoteF__0 = Note{"F__0", 698.46}
	NoteFis0 = Note{"Fis0", 698.46}
	NoteG__0 = Note{"G__0", 783.99}
	NoteGis0 = Note{"Gis0", 830.61}
	// octave 1
	NoteA__1 = Note{"A__1", 880.00}
	NoteAis1 = Note{"Ais1", 932.33}
	NoteB__1 = Note{"B__1", 987.77}
	NoteC__1 = Note{"C__1", 1046.50}
	NoteCis1 = Note{"Cis1", 1108.73}
	NoteD__1 = Note{"D__1", 1174.66}
	NoteDis1 = Note{"Dis1", 1244.51}
	NoteE__1 = Note{"E__1", 1318.51}
	NoteF__1 = Note{"F__1", 1396.91}
	NoteFis1 = Note{"Fis1", 1479.98}
	NoteG__1 = Note{"G__1", 1567.98}
	NoteGis1 = Note{"Gis1", 1661.22}
)

// allNotes contains a slice of notes sorted by semitone from low to high.
var allNotes = []Note{
	NoteA__0, NoteAis0, NoteB__0, NoteC__0, NoteCis0, NoteD__0, NoteDis0, NoteE__0, NoteF__0, NoteFis0, NoteG__0, NoteGis0, NoteA__1, NoteAis1, NoteB__1, NoteC__1, NoteCis1, NoteD__1, NoteDis1, NoteE__1, NoteF__1, NoteFis1, NoteG__1, NoteGis1,
}

// ChordInt defines a handle for potentially non-uniform chord intervals. The name derives from [roman literals] but may
// contain extra symbols for diatonic, chromatic or other chord types.
//
// [roman literals]: https://en.wikipedia.org/wiki/Roman_numeral_analysis
type ChordInt string

const (
	ChordIntI   ChordInt = "I"
	ChordIntII  ChordInt = "II"
	ChordIntIII ChordInt = "III"
	ChordIntIV  ChordInt = "IV"
	ChordIntV   ChordInt = "V"
	ChordIntVI  ChordInt = "VI"
	ChordIntVII ChordInt = "VII"
)

// Note contains ChordKey indicates in which key a chord is supposed to be generated.
type Note struct {
	Name          string
	baseFrequency float64
}

// ChordKey indicates in which key a chord is supposed to be generated.
type ChordKey int

const (
	// ChordKeyMajor indicate a chord's notes towards a major key.
	ChordKeyMajor ChordKey = iota
	// ChordKeyMinor indicate a chord's notes towards a minor key.
	ChordKeyMinor
)

// Chord is a group of harmonic notes played together for their harmonic consonance or dissonance.
type Chord struct {
	Notes []Note
	Key   ChordKey
}

// Length returns the number of notes in this chord. It must not be less than two notes to build a chord.
func (c Chord) Length() int {
	return len(c.Notes)
}

// New creates a note without previously existing notes.
func New() Chord {
	result := Chord{}
	i := intervalChordNotes(NoteC__0, ChordIntI)
	iv := intervalChordNotes(NoteC__0, ChordIntIV)
	v := intervalChordNotes(NoteC__0, ChordIntV)

	result.Notes = []Note{i, iv, v}
	return result
}

func findBaseNoteStart(baseNote Note) int {
	for i, note := range allNotes {
		if note == baseNote {
			return i
		}
	}
	panic(fmt.Sprintf("unable to find base note start for %v", baseNote))
}
