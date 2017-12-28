package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ..#
// #..
// ...

func TestDay22Part1(t *testing.T) {
	testInput := []string{}
	testInput = append(testInput, "..#", "#..", "...")
	actual := VirusGrid(testInput)
	expected := 5587
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}
