package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	)


func TestPart1(t *testing.T) {
	actual := calculateNextDigitSum([]string{"1122"})
	expected := 3
	if expected != actual {
		assert.Equal(t, expected, actual, "first digit (1) matches the second digit and the third digit (2) matches the fourth digit")
	}

	actual = calculateNextDigitSum([]string{"1111"})
	expected = 4
	if expected != actual {
		assert.Equal(t, expected, actual, "each digit (all 1) matches the next")
	}

	actual = calculateNextDigitSum([]string{"1234"})
	expected = 0
	if expected != actual {
		assert.Equal(t, expected, actual, "no digit matches the next")
	}

	actual = calculateNextDigitSum([]string{"91212129"})
	expected = 9
	if expected != actual {
		assert.Equal(t, expected, actual, "only digit that matches the next one is the last digit, 9")
	}
}
//
//Now, instead of considering the next digit, it wants you to consider the digit halfway around the circular list.
//That is, if your list contains 10 items, only include a digit in your sum if the digit 10/2 = 5 steps forward matches it.
//Fortunately, your list has an even number of elements.

func TestPart2(t *testing.T) {
	actual := calculateHalfwaySum([]string{"1212"})
	expected := 6
	if expected != actual {
		assert.Equal(t, expected, actual, "the list contains 4 items, and all four digits match the digit 2 items ahead")
	}

	actual = calculateHalfwaySum([]string{"1221"})
	expected = 0
	if expected != actual {
		assert.Equal(t, expected, actual, "because every comparison is between a 1 and a 2")
	}

	actual = calculateHalfwaySum([]string{"123425"})
	expected = 4
	if expected != actual {
		assert.Equal(t, expected, actual, "because both 2s match each other, but no other digit has a match")
	}

	actual = calculateHalfwaySum([]string{"123123"})
	expected = 12
	if expected != actual {
		assert.Equal(t, expected, actual, "everything has a match")
	}

	actual = calculateHalfwaySum([]string{"12131415"})
	expected = 4
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}