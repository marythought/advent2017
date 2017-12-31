package advent

func HandleJumpInput(input []string) int {
	return findJumps(convertStringArrayToInt(input))
}

func HandleJumpInputPart2(input []string) int {
	return findJumpsPart2(convertStringArrayToInt(input))
}

func findJumps(instructions []int) (jumps int) {
	jumpMap := make(map[int]int)
	for i, val := range instructions {
		_, ok := jumpMap[i]
		if !ok {
			jumpMap[i] = val
		}
	}
	i := 0
	for {
		index := i
		jump, ok := jumpMap[index]
		if !ok {
			break
		}
		i = i + jump
		jumpMap[index]++
		jumps++
	}

	return jumps
}

func findJumpsPart2(instructions []int) (jumps int) {
	jumpMap := make(map[int]int)
	for i, val := range instructions {
		_, ok := jumpMap[i]
		if !ok {
			jumpMap[i] = val
		}
	}
	i := 0
	for {
		index := i
		jump, ok := jumpMap[index]
		if !ok {
			break
		}
		i = i + jump
		if jumpMap[index] >= 3 {
			jumpMap[index]--
		} else {
			jumpMap[index]++
		}
		jumps++
	}

	return jumps
}
