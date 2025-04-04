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

// primary: I, IV, V -> major
// secondary: II, III, VI -> minor
func intervalChordNotes(baseNote Note, desiredInterval ChordInt) Note {
	semitoneInterval, ok := primaryMajorIntervals[desiredInterval]
	if !ok {
		panic(fmt.Sprintf("Invalid desired interval: %s: accepted are roman numerals I..VII", desiredInterval))
	}

	if desiredInterval == "I" {
		return baseNote
	}
	baseNoteStart := findBaseNoteStart(baseNote)
	notesCutToBaseNote := allNotes[baseNoteStart:]

	return notesCutToBaseNote[semitoneInterval]
}
