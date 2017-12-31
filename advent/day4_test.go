package advent

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDay4Part1(t *testing.T) {
	actual := validPassphrase("aa bb cc dd ee")
	expected := true
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual = validPassphrase("aa bb cc dd aa")
	expected = false
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual = validPassphrase("aa bb cc dd aaa")
	expected = true
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}

func TestDay4Part2(t *testing.T) {
	actual := validNoAnagrams("abcde fghij")
	expected := true
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual = validNoAnagrams("abcde xyz ecdab")
	expected = false
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual = validNoAnagrams("a ab abc abd abf abj")
	expected = true
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual = validNoAnagrams("iiii oiii ooii oooi oooo")
	expected = true
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}

	actual = validNoAnagrams("oiii ioii iioi iiio")
	expected = false
	if expected != actual {
		assert.Equal(t, expected, actual, "")
	}
}
