package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9Part1(t *testing.T) {
	actual, gcount := GarbageCollect(`{}`)
	expected := 1
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual, _ = GarbageCollect(`{{{}}}`)
	expected = 6
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual, _ = GarbageCollect(`{{},{}}`)
	expected = 5
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual, _ = GarbageCollect(`{{{},{},{{}}}}`)
	expected = 16
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual, gcount = GarbageCollect(`{<a>,<a>,<a>,<a>}`)
	expected = 1
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual, gcount = GarbageCollect(`{{<ab>},{<ab>},{<ab>},{<ab>}}`)
	expected = 9
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual, gcount = GarbageCollect(`{{<!!>},{<!!>},{<!!>},{<!!>}}`)
	expected = 9
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
	if gcount != 0 {
		assert.Equal(t, expected, gcount, "")
	}

	actual, gcount = GarbageCollect(`{{<a!>},{<a!>},{<a!>},{<ab>}}`)
	expected = 3
	if expected != actual {
		assert.Equal(t, expected, actual, "it should ignore the val after a '!'")
	}
	if gcount != 17 {
		assert.Equal(t, 17, gcount, "")
	}
}
