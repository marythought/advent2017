package advent

import (
	"testing"
	"github.com/stretchr/testify/assert"
	)


func TestPart1(t *testing.T) {
	actual := CalculateNextDigitSum([]string{"1122"})
	expected := 3
	if expected != actual {
		assert.Equal(t, expected, actual, "first digit (1) matches the second digit and the third digit (2) matches the fourth digit")
	}

	actual = CalculateNextDigitSum([]string{"1111"})
	expected = 4
	if expected != actual {
		assert.Equal(t, expected, actual, "each digit (all 1) matches the next")
	}

	actual = CalculateNextDigitSum([]string{"1234"})
	expected = 0
	if expected != actual {
		assert.Equal(t, expected, actual, "no digit matches the next")
	}

	actual = CalculateNextDigitSum([]string{"91212129"})
	expected = 9
	if expected != actual {
		assert.Equal(t, expected, actual, "only digit that matches the next one is the last digit, 9")
	}
}

func TestPart2(t *testing.T) {
	actual := CalculateHalfwaySum([]string{"1212"})
	expected := 6
	if expected != actual {
		assert.Equal(t, expected, actual, "the list contains 4 items, and all four digits match the digit 2 items ahead")
	}

	actual = CalculateHalfwaySum([]string{"1221"})
	expected = 0
	if expected != actual {
		assert.Equal(t, expected, actual, "because every comparison is between a 1 and a 2")
	}

	actual = CalculateHalfwaySum([]string{"123425"})
	expected = 4
	if expected != actual {
		assert.Equal(t, expected, actual, "because both 2s match each other, but no other digit has a match")
	}

	actual = CalculateHalfwaySum([]string{"123123"})
	expected = 12
	if expected != actual {
		assert.Equal(t, expected, actual, "everything has a match")
	}

	actual = CalculateHalfwaySum([]string{"12131415"})
	expected = 4
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}