package advent

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDay6MaxIndex(t *testing.T) {
	max, index := maxIndex([]int{0, 1, 2, 3})
	expectedMax, expectedIndex := 3, 3
	if expectedMax != max {
		assert.Equal(t, max, expectedMax, "")
	}
	if expectedIndex != index {
		assert.Equal(t, index, expectedIndex, "")
	}

	max, index = maxIndex([]int{0, 1, 2, 2, 3, 4, 2, 4})
	expectedMax, expectedIndex = 4, 5
	if expectedMax != max {
		assert.Equal(t, max, expectedMax, "")
	}
	if expectedIndex != index {
		assert.Equal(t, index, expectedIndex, "breaks a tie with lowest index")
	}
}

func TestDay6Part1(t *testing.T) {
	actual := allocate([]int{0, 2, 7, 0})
	expected := 5
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}

func TestDay6Part1WithInput(t *testing.T) {
	var input = []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}
	actual := allocate(input)
	expected := 4074
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}

func TestDay6MakeStringKey(t *testing.T) {
	var array = []int{0, 1, 2, 3}
	actual := makeStringKey(array)
	expected := "0123"
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}

func TestDay6Part2(t *testing.T) {
	var array = []int{2, 4, 1, 2}
	actual := allocate(array)
	expected := 4
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}

func TestDay6Part2WithInput(t *testing.T) {
	// [1 0 14 14 12 12 10 10 8 8 6 6 4 3 2 1]
	var array = []int{1, 0, 14, 14, 12, 12, 10, 10, 8, 8, 6, 6, 4, 3, 2, 1}
	actual := allocate(array)
	expected := 2793
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}
