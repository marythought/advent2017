package advent

import (
	"math"
)

type LocationMap map[coordinate]int

func CalculateDistance(val int) (steps int) {
	lm := MakeLocationMap(val)
	key, _ := mapkey(lm, val)
	x := math.Abs(float64(key.x))
	y := math.Abs(float64(key.y))

	steps = int(x + y)

	return steps
}

func CalculateHighestSum(max int) (result int) {
	return MakeLocationMapWithSums(max)
}

func MakeLocationMapWithSums(max int) (val int) {
	lm := LocationMap{
		{0, 0}: 1,
		{1, 0}: 1,
	}
	x := 1
	y := 0
	direction := "north"
	pivotCount := 1
	stepCount := 0
	for i := 3; i <= max; i++ {
		step(&direction, &x, &y)
		stepCount++
		c := coordinate{x, y}
		val := calculateNeighbors(i, lm, c)
		if val > max {
			return val
		}
		lm[c] = val

		if stepCount == pivotCount {
			pivot(&direction, &pivotCount, &stepCount)
		}
	}
	return
}

// 147  142  133  122   59
// 304    5    4    2   57
// 330   10    1    1   54
// 351   11   23   25   26
// 362  747  806--->   ...
// 265149 is max
func calculateNeighbors(i int, lm LocationMap, c coordinate) (val int) {
	// see if the coordinates around it exist and have values
	coordinates := []coordinate{}
	c1 := coordinate{c.x - 1, c.y - 1}
	c2 := coordinate{c.x - 1, c.y}
	c3 := coordinate{c.x - 1, c.y + 1}
	c4 := coordinate{c.x, c.y - 1}
	c5 := coordinate{c.x, c.y + 1}
	c6 := coordinate{c.x + 1, c.y - 1}
	c7 := coordinate{c.x + 1, c.y}
	c8 := coordinate{c.x + 1, c.y + 1}
	coordinates = append(coordinates, c1, c2, c3, c4, c5, c6, c7, c8)
	for _, coord := range coordinates {
		v, ok := lm[coord]
		if ok {
			val += v
		}
	}
	return val
}

func MakeLocationMap(max int) (lm LocationMap) {
	lm = LocationMap{
		{0, 0}: 1,
		{1, 0}: 2,
	}
	x := 1
	y := 0
	direction := "north"
	pivotCount := 1
	stepCount := 0
	for i := 3; i <= max; i++ {
		step(&direction, &x, &y)
		stepCount++
		c := coordinate{x, y}
		lm[c] = i

		if stepCount == pivotCount {
			pivot(&direction, &pivotCount, &stepCount)
		}
	}
	return lm
}

func step(direction *string, x *int, y *int) {
	switch *direction {
	case "east":
		*x++
	case "north":
		*y++
	case "west":
		*x--
	case "south":
		*y--
	}
	return
}

func pivot(direction *string, pivotCount *int, stepCount *int) {
	switch *direction {
	case "east":
		*direction = "north"
	case "north":
		*direction = "west"
		*pivotCount++
	case "west":
		*direction = "south"
	case "south":
		*direction = "east"
		*pivotCount++
	default:
		*direction = "north"
	}
	*stepCount = 0
	return
}
