package advent

import "strings"

// "pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\nktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\npadx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\njptl (61)\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\ncntj (57)"

func findBase(input string) (base string) {
	lines := splitLines(input)
	return base
}

func splitLines(input string) []string {
	lines := strings.Split(string(input), "\n")
	return lines
}
