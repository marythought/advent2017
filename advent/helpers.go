package advent

import (
	"strconv"
	"strings"
	"sort"
)

func convertStringArrayToInt(input []string)(ret []int){
	for _, v := range input {
		int, _ := strconv.Atoi(v)
		ret = append(ret, int)
	}
	return ret
}

func convertStringToIntSlice(string string) (row []int) {
	nums := strings.Split(string, "\t")
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
	return
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
func minMax(row []int) (int, int) {
	var max = row[0]
	var min = row[0]
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