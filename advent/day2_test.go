package advent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	actual := CalculateCheckSum([]string{"5	1	9	5", "7	5	3", "2	4	6	8"})
	expected := 18
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}

//5 9 2 8
//9 4 7 3
//3 8 6 5
//In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
//In the second row, the two numbers are 9 and 3; the result is 3.
//In the third row, the result is 2.
//In this example, the sum of the results would be 4 + 3 + 2 = 9.

func TestDay2Part2(t *testing.T) {
	actual := CalculateEvenDivideSum([]string{"5	9	2	8", "9	4	7	3", "3	8	6	5"})
	expected := 9
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}
