package audio

import (
	"testing"

	"go-chord-progression/chords"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	chord := chords.NewFor(chords.MinorMelodicIntervals)
	chord.Append(chords.NewFor(chords.MinorMelodicIntervals))
	chord.Append(chords.NewFor(chords.MinorMelodicIntervals))
	chord.Append(chords.NewFor(chords.MinorMelodicIntervals))
	err := Generate(chord)
	require.NoError(t, err)
}
