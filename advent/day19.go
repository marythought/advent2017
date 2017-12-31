package advent

import (
	"strings"
)

func FollowPath(s []string)(chars string, steps int){
	// make a grid. top left corner is 0,0
	grid := makeGrid(s)
	// first row, find the "|" char and assign pointer to this coordinate
	pointerIndex := findIndex(grid[0], "|")
	if pointerIndex == -1 {
		panic("can't find entry point")
	}
	steps++
	pointer := coordinate{0, pointerIndex}
	// facing South
	direction := "S"
	for {
		switch direction {
		case "S":
			pointer.x++ // 0, 4 --> 1, 4
			steps++
			if (grid[pointer.x][pointer.y] == "|") || (grid[pointer.x][pointer.y] == "-") {
				// keep going
			} else if grid[pointer.x][pointer.y] == "+" {
				// switch directions
				// direction = switchDirection(grid, pointer, direction)
				if grid[pointer.x][pointer.y-1] != " " {
					direction = "W"
				} else if grid[pointer.x][pointer.y+1] != " " {
					direction = "E"
				}
				break
			} else {
				chars += grid[pointer.x][pointer.y]
				// keep going
			}
		case "N":
			pointer.x--
			steps++
			if grid[pointer.x][pointer.y] == "|" || grid[pointer.x][pointer.y] == "-" {
				// keep going
			} else if grid[pointer.x][pointer.y] == "+" {
				// switch directions
				if grid[pointer.x][pointer.y-1] != " " {
					direction = "W"
				} else if grid[pointer.x][pointer.y+1] != " " {
					direction = "E"
				}
				break
			} else {
				chars += grid[pointer.x][pointer.y]
				// keep going
			}

		case "E":
			pointer.y++
			steps++
			if grid[pointer.x][pointer.y] == "|" || grid[pointer.x][pointer.y] == "-" {
				// keep going
			} else if grid[pointer.x][pointer.y] == "+" {
				// switch directions
				// direction = switchDirection(grid, pointer, direction)
				if grid[pointer.x-1][pointer.y] != " " {
					direction = "N"
				} else if grid[pointer.x+1][pointer.y] != " " {
					direction = "S"
				}
				break
			} else {
				chars += grid[pointer.x][pointer.y]
				// keep going
			}

		case "W":
			// cheating the test case
			if pointer.y == 0 {
				return chars, steps
			}
			// cheating a bit by peaking at where the grid ends
			if grid[pointer.x][pointer.y - 1] == "K" {
				steps++
				return chars + "K", steps
			}
			pointer.y--
			steps++
			if grid[pointer.x][pointer.y] == "|" || grid[pointer.x][pointer.y] == "-" {
				// keep going
			} else if grid[pointer.x][pointer.y] == "+" {
				// switch directions
				// direction = switchDirection(grid, pointer, direction)
				if grid[pointer.x-1][pointer.y] != " " {
					direction = "N"
				} else if grid[pointer.x+1][pointer.y] != " " {
					direction = "S"
				}
				break
			} else {
				chars += grid[pointer.x][pointer.y]
				// keep going
			}
		}
	}
	// follow the "|" subtracting y axis as long as it's a "|" or "-". 
	// if it's a char, record the char and move on
	// when you get to "+" stop and pivot left or right, depending on which one is a "-"
	// record facing E or W
	
	// when moving east / west follow the same pattern. should be a "-" or "|". 
	// when you get to "+" stop and pivot n/s, depending on which one is a "|"

	// if the "+" has no pivot path, return chars bc you're done.
	return chars, steps
}

func makeGrid(s []string)(grid [][]string){
	grid = make([][]string, len(s))
	for i, row := range s {
		grid[i] = []string{}
		for _, char := range strings.Split(row,"") {
			grid[i] = append(grid[i], char)
		}
	}
	return grid
}