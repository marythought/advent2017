package advent

import (
	"strings"
)

func GarbageCollect(s string) (points, garbageCount int) {
	chars := strings.Split(s, "")
	inGarbage := false
	groups := []string{}
	ignoreNextChar := false
	for _, char := range chars {
		if !ignoreNextChar {
			switch char {
			case "{":
				if !inGarbage {
					groups = append(groups, char)
				} else {
					garbageCount++
				}
			case "}":
				if !inGarbage {
					points += len(groups)
					groups = groups[1:]
				} else {
					garbageCount++
				}
			case "<":
				inGarbage = true
			case ">":
				inGarbage = false
			case "!":
				// ignore next char
				ignoreNextChar = true
			default:
				if inGarbage {
					garbageCount++
				}
			}
		} else {
			ignoreNextChar = false
		}
	}
	return points, garbageCount
}
