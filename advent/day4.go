package advent

import (
	"strings"
	"sort"
)

func CountValidPassphrases(codes []string)(valids int){
	for _, code := range codes {
		if validPassphrase(code) {
			valids++
		}
	}
	return valids
}

func CountValidPassphrasesNoAnagrams(codes []string)(valids int){
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

//Sort String methods via https://stackoverflow.com/a/22698017
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
