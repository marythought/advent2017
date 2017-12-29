package advent

// (0), the initial state before any insertions.
// 0 (1): the spinlock steps forward three times (0, 0, 0), and then inserts the first value, 1, after it. 1 becomes the current position.
// 0 (2) 1: the spinlock steps forward three times (0, 1, 0), and then inserts the second value, 2, after it. 2 becomes the current position.
// 0  2 (3) 1: the spinlock steps forward three times (1, 0, 2), and then inserts the third value, 3, after it. 3 becomes the current position.
// And so on:
//
// 0  2 (4) 3  1
// 0 (5) 2  4  3  1
// 0  5  2  4  3 (6) 1
// 0  5 (7) 2  4  3  6  1
// 0  5  7  2  4  3 (8) 6  1
// 0 (9) 5  7  2  4  3  8  6  1
// Eventually, after 2017 insertions, the section of the circular buffer near the last insertion looks like this:
//
// 1512  1134  151 (2017) 638  1513  851
// Perhaps, if you can identify the value that will ultimately be after the last value written (2017), you can short-circuit the spinlock. In this example, that would be 638.

func Spinlock(step int) int {
	buffer := []int{0}
	index := 0
	for i := 1; i <= 2017; i++ {
		if i == 1 {
			buffer = append(buffer, 1)
			index = 1
		} else {
			index = (index + step) % len(buffer) // 0
			// insert i after the current index position
			newbuffer := make([]int, 2018)
			copy(newbuffer, buffer)
			newbuffer = append(newbuffer[:index+1], i)
			buffer = append(newbuffer, buffer[index+1:]...)
			index++
		}
	}
	return buffer[index+1]
}

func SpinlockPart2(step int) (val int) {
	buffer := []int{0}
	index := 0
	for i := 1; i <= 50000000; i++ {
		if i == 1 {
			buffer = append(buffer, 1)
			index = 1
		} else {
			index = (index + step) % len(buffer) // 0
			// insert i after the current index position
			newbuffer := make([]int, 50000000)
			copy(newbuffer, buffer)
			newbuffer = append(newbuffer[:index+1], i)
			buffer = append(newbuffer, buffer[index+1:]...)
			index++
		}
	}
	next := findIntIndex(buffer, 0)
	return buffer[next+1]
}
