package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10Part1(t *testing.T) {
	// test data
	var testInput = []int{3, 4, 1, 5}
	actual := KnotHash(5, testInput)
	expected := 12
	if expected != actual {
		assert.Equal(t, expected, actual, "list should be 3 4 2 1 [0]")
	}

	var testEvenInput = []int{3, 4, 1, 5, 2}
	actual = KnotHash(6, testEvenInput)
	expected = 0
	if expected != actual {
		assert.Equal(t, expected, actual, "it should work for even number lists")
	}

	// real input
	// 230,1,2,221,97,252,168,169,57,99,0,254,181,255,235,167
	var input = []int{230, 1, 2, 221, 97, 252, 168, 169, 57, 99, 0, 254, 181, 255, 235, 167}
	actual = KnotHash(256, input)
	expected = 2928
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}
