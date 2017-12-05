package advent

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	locationMap := MakeLocationMap(25)
	c1 := coordinate{0,0}
	if locationMap[c1] != 1 {
		assert.Equal(t, c1, locationMap[c1], "grid at 1 should be 0,0")
	}

	c2 := coordinate{1, 0}
	if locationMap[c2] != 2 {
		assert.Equal(t, c2, locationMap[c2], "grid at 2 should be 0,1")
	}

	c4 := coordinate{0, 1}
	if locationMap[c4] != 4 {
		assert.Equal(t, c4, locationMap[c4], "grid at 4 should be 0, 1")
	}

	c5 := coordinate{-1, 1}
	if locationMap[c5] != 5 {
		assert.Equal(t, c5, locationMap[c5], "grid at 5 should be -1, 1")
	}

	c25 := coordinate{2, -2}
	if locationMap[c25] != 25 {
		assert.Equal(t, c25, locationMap[c25], "grid at 25 should be 2, -2")
	}

}