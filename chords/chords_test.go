package chords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intervalChordNotes(t *testing.T) {
	actual := intervalChordNotes(NoteC__0, "I")
	assert.Equal(t, NoteC__0, actual)

	actual = intervalChordNotes(NoteC__0, "IV")
	assert.Equal(t, NoteF__0, actual)

	actual = intervalChordNotes(NoteC__0, "V")
	assert.Equal(t, NoteG__0, actual)
}
