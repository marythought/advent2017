package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay18Part1(t *testing.T) {
	testInput := []string{}
	testInput = append(testInput, "set a 1", "add a 2", "mul a a", "mod a 5", "snd a", "set a 0", "rcv a", "jgz a -1", "set a 1", "jgz a -2")
	actual := Duet(testInput)
	expected := 4
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}
