package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay25Part1(t *testing.T) {
	actual := Turing("test", "A", 6)
	expected := 3
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual = Turing("", "A", 12261543)
	expected = 5744
	if expected != actual {
		assert.Equal(t, expected, actual, "this is the real result")
	}
}
