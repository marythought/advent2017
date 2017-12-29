package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12Part1(t *testing.T) {
	var testInput = "0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5"
	actual := len(FindPipes(testInput, 0))
	expected := 6
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}

func TestDay12Part2(t *testing.T) {
	var testInput = "0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5"
	actual := FindPipeGroups(testInput, 7)
	expected := 2
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}
