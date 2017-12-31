package advent

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDay5Part1(t *testing.T) {
	actual := findJumps([]int{0, 3, 0, 1, -3})
	expected := 5
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}

func TestDay5Part2(t *testing.T) {
	actual := findJumpsPart2([]int{0, 3, 0, 1, -3})
	expected := 10
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}
