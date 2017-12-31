package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay19Part1(t *testing.T) {
	testInput := []string{}
	testInput = append(testInput,
		"    |             ",
		"    |  +--+       ",
		"    A  |  C       ",
		"F---|----E|--+    ",
		"    |  |  |  D    ",
		"    +B-+  +--+    ")
	actual, steps := FollowPath(testInput)
	expected := "ABCDEF"
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	if steps != 38 {
		assert.Equal(t, 38, steps, "")
	}
}
