package advent

import (
	"regexp"
	"strconv"
)

// "pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\nktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\npadx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\njptl (61)\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\ncntj (57)"

type Node struct {
	name         string
	weight       int
	childrenList []string
	children     []Node
}

type nodeMap map[string]Node

var nm nodeMap

func buildTree(input string) {
	nameRegex := regexp.MustCompile(`^\w*`)
	weightRegex := regexp.MustCompile(`\((\d*)\)`)
	childRegex := regexp.MustCompile(`-> (.*)`)

	lines := splitLines(input)
	nm = make(nodeMap)
	for _, node := range lines {
		name := nameRegex.FindString(node)
		weight, _ := strconv.Atoi(weightRegex.FindStringSubmatch(node)[1])
		childrenList := childRegex.FindStringSubmatch(node) //

		n := findOrMakeNode7(name)
		n.weight = weight
		if len(childrenList) > 0 {
			n.childrenList = convertStringListToSlice(childrenList[1])
			addChildren(&n, childrenList)
		}
		nm[name] = n
	}
}

func FindBase(input string) (base string) {
	// who has the most children?
	buildTree(input)
	most := 0
	for k, n := range nm {
		numOfKids := countChildren(n)
		if numOfKids > most {
			base = k
			most = numOfKids
		}
	}
	return base
}

func FindUnbalanced(input string) int {
	// baseName := FindBase(input)
	// base, ok := nm[baseName]
	// if !ok {
	// 	panic("can't find base")
	// }
	// fmt.Println("base is", base.name)
	// for _, n := range base.childrenList {
	// node := nm[n]
	// childw := weighChildren(node)
	// fmt.Println("BASE CHILD")
	// fmt.Println(node.name, node.weight, childw, (node.weight + childw))
	// }
	return 0
}

func findOrMakeNode7(name string) (n Node) {
	n, ok := nm[name]
	if !ok {
		newNode := makeAndAddNode(name)
		n = newNode
	}
	return n
}

func makeAndAddNode(name string) (n Node) {
	n.name = name
	nm[name] = n
	return n
}

func addChildren(n *Node, childrenList []string) {
	if len(childrenList) > 0 {
		children := convertStringListToSlice(childrenList[1])
		for _, child := range children {
			kidNode, ok := nm[child]
			if !ok {
				kidNode = makeAndAddNode(child)
			}
			n.children = append(n.children, kidNode)
		}
		nm[n.name] = *n
	}
}

func countChildren(n Node) (count int) {
	if len(n.children) > 1 {
		for _, child := range n.children {
			if len(child.children) > 1 {
				count += countChildren(child)
			} else {
				count += 1
			}
		}
	}
	return count
}

func weighChildren(n Node) (weight int) {
	if len(n.childrenList) > 0 {
		for _, child := range n.childrenList {
			c := nm[child]
			if len(c.childrenList) > 0 {
				weight += weighChildren(c)
			} else {
				weight += c.weight
			}
			// fmt.Println("my name is", c.name, "I weigh", c.weight, "my parent is", n.name)
		}
	}
	return weight
}
