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
