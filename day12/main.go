package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/matgd/advent2021/utils"
)

const (
	start string = "start"
	end          = "end"
)

type node struct {
	name       string
	isSmall    bool
	visited    bool
	neighbours []*node
}

func (n node) String() string {
	nNames := make([]string, 0, len(n.neighbours))
	for _, neighbour := range n.neighbours {
		nNames = append(nNames, neighbour.name)
	}

	return fmt.Sprintf("Name: %v, Neighbours: %v\n", n.name, nNames)
}

func (n *node) isNeighbour(otherNode *node) bool {
	for _, neighbour := range n.neighbours {
		if neighbour == otherNode {
			return true
		}
	}
	return false
}

type graph struct {
	nodes          []*node
	nodesAddresses map[string]*node
	start          *node
	end            *node
}

func (g graph) String() string {
	return fmt.Sprintf(
		"Nodes: [%v],\nAddrs: {%v},\nStart: <%v>,\nEnd: <%v>\n",
		g.nodes, g.nodesAddresses, g.start, g.end,
	)
}

func (g *graph) DFS(n *node, acc *[]*node, allPaths []*[]*node) {
	n.visited = true
	*acc = append(*acc, n)

	var accCopy []*node
	copy(accCopy, *acc)
	fmt.Println(">>>>>>>>>>", *acc, accCopy)
	allPaths = append(allPaths, &accCopy)
	fmt.Println(">>>", allPaths[0])

	if n.name == end {
		return
	}
	for _, nb := range n.neighbours {
		if !nb.visited || !nb.isSmall {
			g.DFS(nb, acc, allPaths)
		}
	}
}

func loadGraph(inputLines []string) graph {
	g := graph{
		nodes:          make([]*node, 0, len(inputLines)),
		nodesAddresses: map[string]*node{},
		start:          nil,
		end:            nil,
	}

	for _, line := range inputLines {
		split := strings.Split(line, "-")

		n1 := new(node)
		n2 := new(node)
		if _, exists := g.nodesAddresses[split[0]]; !exists {
			n1SmallCave := true
			if strings.ToUpper(split[0]) == split[0] {
				n1SmallCave = false
			}
			n1 = &node{
				name:       split[0],
				isSmall:    n1SmallCave,
				visited:    false,
				neighbours: make([]*node, 0, len(inputLines)),
			}

			g.nodes = append(g.nodes, n1)
			g.nodesAddresses[split[0]] = n1
			if split[0] == start {
				g.start = n1
			} else if split[0] == end {
				g.end = n1
			}
		} else {
			n1 = &*g.nodesAddresses[split[0]]
		}

		if _, exists := g.nodesAddresses[split[1]]; !exists {
			n2SmallCave := true
			if strings.ToUpper(split[1]) == split[1] {
				n2SmallCave = false
			}
			n2 = &node{
				name:       split[1],
				isSmall:    n2SmallCave,
				visited:    false,
				neighbours: make([]*node, 0, len(inputLines)),
			}

			g.nodes = append(g.nodes, n2)
			g.nodesAddresses[split[1]] = n2
			if split[1] == start {
				g.start = n2
			} else if split[1] == end {
				g.end = n2
			}
		} else {
			n2 = &*g.nodesAddresses[split[1]]
		}

		if !(n1.isNeighbour(n2)) {
			n1.neighbours = append(n1.neighbours, n2)
			n2.neighbours = append(n2.neighbours, n1)
		}
	}
	return g
}

func task1(inputLines []string) int {
	g := loadGraph(inputLines)

	acc := make([]*node, 0, 1000)
	allPaths := make([]*[]*node, 0, 1000)

	g.DFS(g.start, &acc, allPaths)
	fmt.Println(acc)
	fmt.Println(allPaths)
	fmt.Printf("\n\n\n")
	log.Fatal("BORKED!")
	return 1
}

func main() {
	input := utils.GetStringsFromInputFile("input_example.txt")
	task1(input)
}
