package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	locationMap := MakeLocationMap(25)
	c1 := coordinate{0, 0}
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

// 147  142  133  122   59
// 304    5    4    2   57
// 330   10    1    1   54
// 351   11   23   25   26
// 362  747  806--->   ...
// 265149 is max
func TestDay3Part2(t *testing.T) {
	actual := MakeLocationMapWithSums(25)
	expected := 26
	if expected != actual {
		assert.Equal(t, expected, actual, "highest result should be 26")
	}
}
