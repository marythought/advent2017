package advent

import "math"

type LocationMap map[coordinate]int
type coordinate struct {
	x, y int
}

func CalculateDistance(val int)(steps int) {
	lm := MakeLocationMap(val)
	key, _ := mapkey(lm, val)
	x := math.Abs(float64(key.x))
	y := math.Abs(float64(key.y))

	steps = int(x + y)

	return steps
}

func MakeLocationMap(max int)(lm LocationMap){
	lm = LocationMap{
		{0,0}: 1,
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

func step(direction *string, x *int, y *int)(){
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

func pivot(direction *string, pivotCount *int, stepCount *int)(){
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
