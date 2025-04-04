package chords

import (
	"fmt"
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

func Test_findBaseNoteStart(t *testing.T) {
	t.Run("should return the semitone index for a given note", func(t *testing.T) {
		type args struct {
			baseNote Note
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			{"A", args{NoteA__0}, 0},
			{"C", args{NoteC__0}, 3},
			{"Gis", args{NoteGis0}, 11},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equalf(t, tt.want, findBaseNoteStart(tt.args.baseNote), "findBaseNoteStart(%v)", tt.args.baseNote)
			})
		}
	})

	t.Run("should panic for invalid input", func(t *testing.T) {
		// given
		invalidInput := Note{"asdf", 0}

		defer func() {
			// then
			if r := recover(); r != nil {
				fmt.Println("successfully panicked", r)
				assert.Contains(t, r, "unable to find base note start for")
			}
		}()

		// when
		findBaseNoteStart(invalidInput)

		// not-then
		assert.Fail(t, "should panic for invalid input")
	})
}
