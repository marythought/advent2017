package advent

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func FindPipes(input string, startingPipe int) []int {
	pipeMap := makePipeMap(input)
	toProcess := []int{startingPipe}
	connected := []int{}
	processed := []int{}
	for {
		if len(toProcess) < 1 {
			break
		}
		pipe := toProcess[0]
		v, ok := pipeMap[pipe]
		if !ok {
			break
		}
		for _, num := range v {
			if findIntIndex(connected, num) < 0 {
				connected = append(connected, num)
			}
			if findIntIndex(processed, num) < 0 && findIntIndex(toProcess, num) < 0 {
				toProcess = append(toProcess, num)
			}
		}
		processed = append(processed, pipe)
		if len(toProcess) > 0 {
			toProcess = toProcess[1:]
		}
	}
	return connected
}

func FindPipeGroups(input string, max int) (groups int) {
	connected := FindPipes(input, 0)
	// 0 through 1999 are in the neighborhood.
	groups++
	for i := 1; i < max; i++ {
		if findIntIndex(connected, i) < 0 {
			group := FindPipes(input, i)
			groups++
			connected = append(connected, group...)
		}
	}
	return groups
}

func makePipeMap(input string) map[int][]int {
	slice := strings.Split(input, "\n")
	regex := regexp.MustCompile(`(\d+) <-> (\d+,*.*)`)
	pipeMap := make(map[int][]int)
	for _, v := range slice {
		pipeData := regex.FindStringSubmatch(v)
		start, _ := strconv.Atoi(pipeData[1])
		links := pipeData[2]
		intLinks := []int{}
		for _, v := range strings.Split(links, ", ") {
			strings.TrimSpace(v)
			link, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("can't convert string to int in the pipe links")
			}
			intLinks = append(intLinks, link)
		}
		pipeMap[start] = intLinks
	}
	return pipeMap
}
