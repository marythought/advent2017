package advent

type keyVal struct {
	value int
	move  string
	next  string
}

type tickerTape []int

func Turing(s string, begin string, checks int) (checksum int) {
	var key map[string][]keyVal
	if s == "test" {
		key = makeTestKey()
	} else {
		key = makeRealKey()
	}

	tape := make(tickerTape, 101)
	counter := len(tape) / 2
	state := begin
	for i := 0; i < checks; i++ {
		val := key[state][tape[counter]].value
		move := key[state][tape[counter]].move
		next := key[state][tape[counter]].next

		// set value
		tape[counter] = val
		// check if counter is out of index, if yes add 10
		if move == "R" {
			if counter+1 > len(tape) {
				tape = append(tape, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}...)
			}
			// move counter
			counter++
		} else {
			if counter == 0 {
				tape = append([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, tape...)
				// reset the index
				counter = 10
			}
			counter--
		}
		state = next
	}
	// Perform a diagnostic checksum after x checks.
	for _, v := range tape {
		checksum += v
	}

	return checksum
}

func makeTestKey() (key map[string][]keyVal) {
	key = make(map[string][]keyVal)
	key["A"] = append(key["A"], keyVal{1, "R", "B"}, keyVal{0, "L", "B"})
	key["B"] = append(key["B"], keyVal{1, "L", "A"}, keyVal{1, "R", "A"})
	return key
}

// Begin in state A.
// Perform a diagnostic checksum after 12261543 steps.
func makeRealKey() (key map[string][]keyVal) {
	key = make(map[string][]keyVal)
	key["A"] = append(key["A"], keyVal{1, "R", "B"}, keyVal{0, "L", "C"})
	key["B"] = append(key["B"], keyVal{1, "L", "A"}, keyVal{1, "R", "C"})
	key["C"] = append(key["C"], keyVal{1, "R", "A"}, keyVal{0, "L", "D"})
	key["D"] = append(key["D"], keyVal{1, "L", "E"}, keyVal{1, "L", "C"})
	key["E"] = append(key["E"], keyVal{1, "R", "F"}, keyVal{1, "R", "A"})
	key["F"] = append(key["F"], keyVal{1, "R", "A"}, keyVal{1, "R", "E"})
	return key
}
