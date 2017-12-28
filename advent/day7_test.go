package advent

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

var testInput7 string = "pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\nktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\npadx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\njptl (41) -> abcd, defg\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\ncntj (57)\nabcd (10)\ndefg (10)"

// func TestDay7Part1(t *testing.T) {
// 	actual := FindBase(testInput7)
// 	expected := "tknk"
// 	if expected != actual {
// 		assert.Equal(t, actual, expected, "")
// 	}
// }

// ugml + (gyxo + ebii + jptl) = 68 + (61 + 61 + (jptl = 41 + abcd(10) + defg(10)))) = 251
// padx + (pbga + havc + qoyq) = 45 + (66 + 66 + 66) = 243
// fwft + (ktlj + cntj + xhth) = 72 + (57 + 57 + 57) = 243

// func TestDay7Part2(t *testing.T) {
// 	FindUnbalanced(testInput7, "")
// }
//
// func TestDay7Part2RealInput(t *testing.T) {
// 	input, err := ioutil.ReadFile("../input/day7.txt")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	FindUnbalanced(string(input), "apktiv")
// }

func TestConvertStringSlice(t *testing.T) {
	actual := convertStringListToSlice("mary, josh, aurelio")
	expected := []string{"mary", "josh", "aurelio"}
	if expected != nil {
		assert.Equal(t, actual, expected, "")
	}
	actual = convertStringListToSlice("mary,josh,aurelio")
	expected = []string{"mary", "josh", "aurelio"}
	if expected != nil {
		assert.Equal(t, actual, expected, "")
	}
}

func TestSplitLines(t *testing.T) {
	actual := splitLines(testInput7)
	expected := "pbga (66)"
	if actual[0] != expected {
		assert.Equal(t, actual[0], expected, "")
	}
}
