package advent

import (
	"regexp"
	"strconv"
)

type Node8 struct {
	name  string
	value int
}

type nodeMap8 map[string]Node8

type Instruction struct {
	originNode      string
	operator        string
	amount          int
	conditionalNode string
	conditional     []string
}

var nm8 nodeMap8
var regex *regexp.Regexp
var highest int

func FindHighestValEver(lines []string) int {
	instructions := buildNodesAndInstructions(lines)
	for _, i := range instructions {
		executeInstruction(i)
	}
	return highest
}

func FindLargestValue(lines []string) int {
	instructions := buildNodesAndInstructions(lines)
	for _, i := range instructions {
		executeInstruction(i)
	}
	return findLargestValueInNodeMap8().value
}

func buildNodesAndInstructions(lines []string) (instructions []Instruction) {
	regex = regexp.MustCompile(`(\w*)\s(inc|dec)\s(-?\d*)\sif\s(\w*)\s(\S*)\s(-?\d*)`)
	nm8 = make(nodeMap8)

	for _, line := range lines {
		lineQualities := regex.FindStringSubmatch(line)
		i := makeInstruction(lineQualities)
		instructions = append(instructions, i)
	}
	return instructions
}

func findOrMakeNode(name string) (n Node8) {
	n, ok := nm8[name]
	if !ok {
		newNode8 := makeDay8Node(name)
		n = newNode8
	}
	return n
}

func makeDay8Node(name string) (n Node8) {
	n.name = name
	nm8[name] = n
	return n
}

func makeInstruction(lineQualities []string) (i Instruction) {
	amount, _ := strconv.Atoi(lineQualities[3])
	return Instruction{
		originNode:      lineQualities[1],
		operator:        lineQualities[2],
		amount:          amount,
		conditionalNode: lineQualities[4],
		conditional:     []string{lineQualities[5], lineQualities[6]},
	}
}

func executeInstruction(i Instruction) {
	if conditionalPass(i) {
		n := findOrMakeNode(i.originNode)
		amt := i.amount
		switch i.operator {
		case "inc":
			n.value += amt
		case "dec":
			n.value -= amt
		}
		if n.value > highest {
			highest = n.value
		}
		nm8[n.name] = n
	}
	return
}

func conditionalPass(i Instruction) bool {
	n := findOrMakeNode(i.conditionalNode)
	check := i.conditional
	checknum, _ := strconv.Atoi(check[1])
	switch check[0] {
	case ">":
		return n.value > checknum
	case ">=":
		return n.value >= checknum
	case "<":
		return n.value < checknum
	case "<=":
		return n.value <= checknum
	case "==":
		return n.value == checknum
	case "!=":
		return n.value != checknum
	}
	return true
}

func findLargestValueInNodeMap8() (n Node8) {
	largestAmt := 0
	for _, v := range nm8 {
		if v.value > largestAmt {
			largestAmt = v.value
			n = v
		}
	}
	return n
}
