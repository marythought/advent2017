package advent

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

var testInput string = "pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\nktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\npadx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\njptl (61)\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\ncntj (57)"

func TestDay7Part1(t *testing.T) {
	actual := findBase(testInput)
	expected := "tknk"
	if expected != actual {
		assert.Equal(t, actual, expected, "")
	}
}

func TestSplitLines(t *testing.T) {
	actual := splitLines(testInput)
	expected := "pbga (66)"
	if actual[0] != expected {
		assert.Equal(t, actual[0], expected, "")
	}
}
