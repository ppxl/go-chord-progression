package chords

import "fmt"

// primaryMajorIntervals counts the actual semitones between the primary chords, making it possible to mathematically
// count primaries independently of the base note.
// I II III IV V VI VII I¹
// 0, 2, 2, 1, 2, 2, 2, 1
var primaryMajorIntervals = map[ChordInt]int{
	ChordIntI:   0,
	ChordIntII:  0 + 2,
	ChordIntIII: 0 + 2 + 2,
	ChordIntIV:  0 + 2 + 2 + 1,
	ChordIntV:   0 + 2 + 2 + 1 + 2,
	ChordIntVI:  0 + 2 + 2 + 1 + 2 + 2,
	ChordIntVII: 0 + 2 + 2 + 1 + 2 + 2 + 2,
}

// primary chords are: I, IV, V.
// In a major key, all primary chords are major triads.
// In a minor key, the primary chords I and IV are minor triads, V is still a major triad.
//
// secondary: II, III, VI
// In a major key, all secondary chords are minor triads.
// In a minor key, the primary chords III and VI are major triads, II is a diminished triad.
func intervalChordNotes(baseNote Note, desiredInterval ChordInt) Note {
	semitoneInterval, ok := primaryMajorIntervals[desiredInterval]
	if !ok {
		panic(fmt.Sprintf("unsupported interval: %s: accepted are roman numerals I..VII", desiredInterval))
	}

	if desiredInterval == ChordIntI {
		return baseNote
	}
	// find the base note in all notes and then count the semitones up to get the desired interval note
	// f. i. IV (5 semitones up) based on C: A,A#,B,|start here>C,C#,D,D#,E,F<---found
	baseNoteStart := findBaseNoteStart(baseNote)
	return allNotes[baseNoteStart+semitoneInterval]
}
