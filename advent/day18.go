package advent

import (
	"regexp"
	"strconv"
	"strings"
)

func Duet(input []string) (soundValue int) {
	tablet := make(map[string]int)
	regex := regexp.MustCompile(`^(\w{3}) (\w)(.*)`)
	counter := 0
	for {
		var y string
		var setChar string
		if counter > len(input) || counter < 0 {
			break
		}
		line := input[counter]
		directions := regex.FindStringSubmatch(line)
		command := directions[1]
		x := directions[2]
		if len(directions) > 3 {
			y = strings.TrimSpace(directions[3])
		}
		switch command {
		case "set":
			// set X Y sets register X to the value of Y, which can be an int or a key
			val, err := strconv.Atoi(y)
			if err != nil {
				setChar = y
			}
			if setChar != "" {
				v, ok := tablet[setChar]
				if !ok {
					tablet[setChar] = 0
					tablet[x] = 0
				} else {
					tablet[x] = v
				}
			} else {
				tablet[x] = val
			}
			counter++
		case "snd":
			// snd X plays a sound with a frequency equal to the value of X, which is a key
			_, ok := tablet[x]
			if !ok {
				tablet[x] = 0
			}
			soundValue = tablet[x]
			counter++
		case "add":
			// add X Y increases register X by the value of Y, can be int or key
			val, err := strconv.Atoi(y)
			if err != nil {
				setChar = y
			}
			if setChar != "" {
				v, ok := tablet[setChar]
				if !ok {
					tablet[setChar] = 0
				} else {
					reg, ok := tablet[x]
					if !ok {
						tablet[x] = val
					}
					tablet[x] = reg + v
				}
			} else {
				reg, ok := tablet[x]
				if !ok {
					tablet[x] = val
				}
				tablet[x] = reg + val
			}
			counter++
		case "mul":
			// mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y. int or key.
			val, err := strconv.Atoi(y)
			if err != nil {
				setChar = y
			}
			if setChar != "" {
				v, ok := tablet[setChar]
				if !ok {
					tablet[setChar] = 0
				} else {
					reg, ok := tablet[x]
					if !ok {
						tablet[x] = 0
					}
					tablet[x] = reg * v
				}
			} else {
				reg, ok := tablet[x]
				if !ok {
					tablet[x] = 0
				}
				tablet[x] = reg * val
			}
			counter++
		case "mod":
			// mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y (that is, it sets X to the result of X modulo Y).
			val, err := strconv.Atoi(y)
			if err != nil {
				setChar = y
			}
			if setChar != "" {
				v, ok := tablet[setChar]
				if !ok {
					tablet[setChar] = 0
				} else {
					reg, ok := tablet[x]
					if !ok {
						tablet[x] = 0
					}
					tablet[x] = reg % v
				}
			} else {
				reg, ok := tablet[x]
				if !ok {
					tablet[x] = 0
				}
				tablet[x] = reg % val
			}
			counter++
		case "rcv":
			// rcv X recovers the frequency of the last sound played, but only when the value of X is not zero. (If it is zero, the command does nothing.)
			if soundValue > 0 {
				return soundValue
			}
			counter++
		case "jgz":
			// jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
			reg, ok := tablet[x]
			if !ok {
				tablet[x] = 0
			}
			if reg > 0 {
				val, err := strconv.Atoi(y)
				if err != nil {
					setChar = y
				}
				if setChar != "" {
					v, ok := tablet[setChar]
					if !ok {
						tablet[setChar] = 0
					} else {
						counter = counter + v
					}
				} else {
					counter = counter + val
				}
			} else {
				counter++
			}
		}
	}
	return soundValue
}
