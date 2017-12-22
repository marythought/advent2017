package advent

import "fmt"

//
// To achieve this, begin with a list of numbers from 0 to 255, a current position which begins at 0 (the first element in the list), a skip size (which starts at 0), and a sequence of lengths (your puzzle input). Then, for each length:
//
// Reverse the order of that length of elements in the list, starting with the element at the current position.
// Move the current position forward by that length plus the skip size.
// Increase the skip size by one.

func populateList(size int) (list []int) {
	for i := 0; i < size; i++ {
		list = append(list, i)
	}
	return list
}

func KnotHash(size int, input []int) int {
	list := populateList(size)
	position := 0
	skipSize := 0
	for index, length := range input {
		fmt.Println("loop number ", index, "input is ", length, "index is", position, "list is", list)
		startIndex := position
		endIndex := (startIndex + length) % size
		if length < 2 {
		} else if endIndex > startIndex {
			loop := list[startIndex:endIndex]
			reverseSlice(loop)
			counter := 0
			for i, _ := range list[startIndex:endIndex] {
				list[i] = loop[counter]
				counter++
			}
		} else {
			loop := append(list[startIndex:], list[0:endIndex]...)
			reverseSlice(loop)

			counter := 0
			for i := startIndex; i < size; i++ {
				list[i] = loop[counter]
				counter++
			}
			for i := 0; i < endIndex; i++ {
				list[i] = loop[counter]
				counter++
			}
		}
		position = (position + length + skipSize) % size
		skipSize++
	}
	return list[0] * list[1]
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
