package advent

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

var testInput8 string = "b inc 5 if a > 1\na inc 1 if b < 5\nc dec -10 if a >= 1\nc inc -20 if c == 10"

func TestDay8Part1(t *testing.T) {
	actual := FindLargestValue(splitLines(testInput8))
	expected := 1
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}

func TestDay8Part2(t *testing.T) {
	actual := FindHighestValEver(splitLines(testInput8))
	expected := 10
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}
