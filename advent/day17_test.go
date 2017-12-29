package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17Part1(t *testing.T) {
	actual := Spinlock(3)
	expected := 638
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}
