package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16Part1(t *testing.T) {
	var testInput = "s1,x3/4,pe/b"

	actual := DancingPrograms(testInput, "abcde")
	expected := "baedc"
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}

func TestDay16Part2(t *testing.T) {
	var testInput = "s1,x3/4,pe/b"

	actual, loops := findLoop(testInput, "abcde", 2)
	expected := "ceadb"
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	if 2 != loops {
		assert.Equal(t, 2, loops, "if it completes it should return that number of loops")
	}

	actual, loops = findLoop(testInput, "abcde", 4)
	if "abcde" != actual {
		assert.Equal(t, "abcde", actual, "")
	}
	if 4 != loops {
		assert.Equal(t, 4, loops, "it returns the original code in 4 loops", actual)
	}
}
