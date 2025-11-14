package chords

import "fmt"

// MajorIntervals counts the actual semitones between the major chords, making it possible to mathematically
// count primaries, etc. independently of the base note.
// I II III IV V VI VII I¹
// 0, 2, 2, 1, 2, 2, 2, 1
var MajorIntervals = map[ChordInt]int{
	ChordIntI:   0,
	ChordIntII:  0 + 2,
	ChordIntIII: 0 + 2 + 2,
	ChordIntIV:  0 + 2 + 2 + 1,
	ChordIntV:   0 + 2 + 2 + 1 + 2,
	ChordIntVI:  0 + 2 + 2 + 1 + 2 + 2,
	ChordIntVII: 0 + 2 + 2 + 1 + 2 + 2 + 2,
}

// minorHarmonicIntervals counts the actual semitones between the harmonic minor chords, making it possible to mathematically
// count primaries, etc. independently of the base note.
// I II III IV V VI VII I¹
// 0, 2, 1, 2, 2, 1, 3, 1
var minorHarmonicIntervals = map[ChordInt]int{
	ChordIntI:   0,
	ChordIntII:  0 + 2,
	ChordIntIII: 0 + 2 + 1,
	ChordIntIV:  0 + 2 + 1 + 2,
	ChordIntV:   0 + 2 + 1 + 2 + 2,
	ChordIntVI:  0 + 2 + 1 + 2 + 2 + 1,
	ChordIntVII: 0 + 2 + 1 + 2 + 2 + 1 + 3,
}

// minorNaturalIntervals counts the actual semitones between the natural minor chords, making it possible to mathematically
// count primaries, etc. independently of the base note.
// I II III IV V VI VII I¹
// 0, 2, 1, 2, 2, 1, 2, 2
var minorNaturalIntervals = map[ChordInt]int{
	ChordIntI:   0,
	ChordIntII:  0 + 2,
	ChordIntIII: 0 + 2 + 1,
	ChordIntIV:  0 + 2 + 1 + 2,
	ChordIntV:   0 + 2 + 1 + 2 + 2,
	ChordIntVI:  0 + 2 + 1 + 2 + 2 + 1,
	ChordIntVII: 0 + 2 + 1 + 2 + 2 + 1 + 2,
}

// MinorMelodicIntervals counts the actual semitones between the melodic minor chords, making it possible to mathematically
// count primaries, etc. independently of the base note.
// I II III IV V VI VII I¹
// 0, 2, 1, 2, 2, 2, 2, 1
var MinorMelodicIntervals = map[ChordInt]int{
	ChordIntI:   0,
	ChordIntII:  0 + 2,
	ChordIntIII: 0 + 2 + 1,
	ChordIntIV:  0 + 2 + 1 + 2,
	ChordIntV:   0 + 2 + 1 + 2 + 2,
	ChordIntVI:  0 + 2 + 1 + 2 + 2 + 2,
	ChordIntVII: 0 + 2 + 1 + 2 + 2 + 2 + 2,
}

type Scale map[ChordInt]int

// primary chords are: I, IV, V.
// In a major key, all primary chords are major triads.
// In a minor key, the primary chords I and IV are minor triads, V is still a major triad.
//
// secondary: II, III, VI
// In a major key, all secondary chords are minor triads.
// In a minor key, the primary chords III and VI are major triads, II is a diminished triad.
func intervalChordNotes(baseNote Note, desiredInterval ChordInt) Note {
	return intervalChordNotesByScale(baseNote, desiredInterval, MajorIntervals)
}
func intervalChordNotesByScale(baseNote Note, desiredInterval ChordInt, scale Scale) Note {
	semitoneInterval, ok := scale[desiredInterval]
	if !ok {
		panic(fmt.Sprintf("unsupported scale: %s: accepted are roman numerals I..VII", desiredInterval))
	}

	if desiredInterval == ChordIntI {
		return baseNote
	}
	// find the base note in all notes and then count the semitones up to get the desired scale note
	// f. i. IV (5 semitones up) based on C: A,A#,B,|start here>C,C#,D,D#,E,F<---found
	baseNoteStart := findBaseNoteStart(baseNote)
	return allNotes[baseNoteStart+semitoneInterval]
}
