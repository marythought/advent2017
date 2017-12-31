package advent

import (
	"strings"
)

func CountValidPassphrases(codes []string) (valids int) {
	for _, code := range codes {
		if validPassphrase(code) {
			valids++
		}
	}
	return valids
}

func CountValidPassphrasesNoAnagrams(codes []string) (valids int) {
	for _, code := range codes {
		if validNoAnagrams(code) {
			valids++
		}
	}
	return valids
}

func validPassphrase(code string) bool {
	unique := map[string]bool{}
	chunks := strings.Split(code, " ")
	for _, chunk := range chunks {
		if unique[chunk] {
			return false
		} else {
			unique[chunk] = true
		}
	}
	return true
}

func validNoAnagrams(code string) bool {
	unique := map[string]bool{}
	chunks := strings.Split(code, " ")
	for _, chunk := range chunks {
		s := SortString(chunk)
		if unique[s] {
			return false
		} else {
			unique[s] = true
		}
	}
	return true
}
