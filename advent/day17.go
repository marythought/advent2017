package advent

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
