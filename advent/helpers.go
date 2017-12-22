package advent

import (
	"sort"
	"strconv"
	"strings"
)

func convertStringArrayToInt(input []string) (ret []int) {
	for _, v := range input {
		int, _ := strconv.Atoi(v)
		ret = append(ret, int)
	}
	return ret
}

func convertStringListToSlice(input string) (cleanArray []string) {
	a := strings.Split(input, ",")
	for _, s := range a {
		cleanArray = append(cleanArray, strings.TrimSpace(s))
	}
	return cleanArray
}

func convertStringToIntSlice(s string) (row []int) {
	nums := strings.Split(s, "\t")
	for _, n := range nums {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		row = append(row, num)
	}
	return row
}

func mapkey(m LocationMap, value int) (key coordinate, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return key, ok
}

func splitLines(input string) []string {
	lines := strings.Split(string(input), "\n")
	return RemoveTrailingEmptyStringsInStringArray(lines)
}

// from https://siongui.github.io/2017/01/19/go-remove-leading-and-trailing-empty-strings-in-string-slice/
func RemoveTrailingEmptyStringsInStringArray(sa []string) []string {
	lastNonEmptyStringIndex := len(sa) - 1
	for i := lastNonEmptyStringIndex; i >= 0; i-- {
		if sa[i] == "" {
			lastNonEmptyStringIndex--
		} else {
			break
		}
	}
	return sa[0 : lastNonEmptyStringIndex+1]
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

//via https://stackoverflow.com/a/45976758
func minMax(row []int) (min int, max int) {
	max = row[0]
	min = row[0]
	for _, value := range row {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func maxIndex(array []int) (max int, index int) {
	max = array[0]
	index = 0
	for i, value := range array {
		if max < value {
			max = value
			index = i
		}
	}
	return max, index
}
