package advent

import (
	"strings"
)

type carrier struct {
	direction   string
	coordinates [2]int
	awake       bool
}

func CleanInput(list []string) (grid [][]string) {
	for _, row := range list {
		grid = append(grid, strings.Split(row, ""))
	}
	return grid
}

func VirusGrid(list []string) (infectedCount int) {
	grid := CleanInput(list)
	grid = expand(grid)
	bursts := 10000
	compass := []string{"N", "E", "S", "W"}
	startx := len(grid) / 2
	starty := len(grid[startx]) / 2
	c := carrier{"N", [2]int{startx, starty}, false}
	for i := 0; i < bursts; i++ {
		dirIndex := findIndex(compass, c.direction)
		if grid[c.coordinates[0]][c.coordinates[1]] == "#" {
			c.direction = compass[(dirIndex+1)%len(compass)]
			grid[c.coordinates[0]][c.coordinates[1]] = "."
		} else {
			if dirIndex == 0 {
				c.direction = compass[3]
			} else {
				c.direction = compass[(dirIndex - 1)]
			}
			grid[c.coordinates[0]][c.coordinates[1]] = "#"
			infectedCount++
		}
		// move 1 step forward
		switch c.direction {
		case "N":
			c.coordinates[0] = c.coordinates[0] - 1
		case "E":
			c.coordinates[1] = c.coordinates[1] + 1
		case "W":
			c.coordinates[1] = c.coordinates[1] - 1
		case "S":
			c.coordinates[0] = c.coordinates[0] + 1
		}
	}
	return infectedCount
}

func expand(grid [][]string) [][]string {
	max := 1000
	for index, row := range grid {
		for i := 0; i < max; i++ {
			row = append(row, ".")
		}
		for i := 0; i < max; i++ {
			row = append([]string{"."}, row...)
		}
		grid[index] = row
	}

	//add max empty rows at the end
	for i := 0; i < max; i++ {
		row := []string{}
		for j := 0; j < 2025; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	//add max empty rows at the beginning
	start := [][]string{}
	for i := 0; i < max; i++ {
		row := []string{}
		for j := 0; j < 2025; j++ {
			row = append(row, ".")
		}
		start = append(start, row)
	}

	grid = append(start, grid...)
	return grid
}
