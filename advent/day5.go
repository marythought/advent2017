package advent

import (
	"strconv"
)

func HandleJumpInput(input []string)(int){
	instructions := []int{}
	for _, v := range input {
		int, _ := strconv.Atoi(v)
		instructions = append(instructions, int)
	}
	return findJumps(instructions)
}

func HandleJumpInputPart2(input []string)(int){
	instructions := []int{}
	for _, v := range input {
		int, _ := strconv.Atoi(v)
		instructions = append(instructions, int)
	}
	return findJumpsPart2(instructions)
}

func findJumps(instructions []int)(jumps int){
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
		jumpMap[index] = jumpMap[index] + 1
		jumps++
	}

	return jumps
}

func findJumpsPart2(instructions []int)(jumps int){
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
			jumpMap[index] = jumpMap[index] - 1
		} else {
			jumpMap[index] = jumpMap[index] + 1
		}
		jumps++
	}

	return jumps
}



