package aoc6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const NOT_FOUND = 999999
type Node struct {
	value string
}

type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	if g.nodes == nil {
		g.nodes = []*Node{}
	}
	g.nodes = append(g.nodes, n)
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
}

func (graph *ItemGraph) Print() string {
	str := ""
	for node_ind, node_val := range graph.nodes {
		for _, edgeValue := range graph.edges[*graph.nodes[node_ind]]{
			str += fmt.Sprintf("%v -> %v \n", node_val.value, edgeValue.value)
		}
	}
	return str
}

func LoadLinkFile(FileToLoad string) ItemGraph {
	graph := ItemGraph{
		nodes: nil,
		edges: nil,
	}

	file, err := os.Open(FileToLoad)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ")")
		Node1 := graph.addNodeIfDoesNotExist(split[0])
		Node2 := graph.addNodeIfDoesNotExist(split[1])
		graph.AddEdge(Node2,Node1)
	}

	return graph
}

func GetLinksNumber(FileToLoad string) int {
	graph := LoadLinkFile(FileToLoad)
	links := 0
	for _, node1_obj := range graph.nodes{
		for _, node2_obj := range graph.nodes{
			if graph.canReach(node1_obj,node2_obj) && node2_obj != node1_obj {
				links += 1
			}
		}
	}
	return links
}

func FindDistanceBetween(FileToLoad string,start string,end string) int {
	graph := LoadLinkFile(FileToLoad)
	startNode := graph.edges[*graph.getNode(start)][0]
	endNode := graph.edges[*graph.getNode(end)][0]

	distance := graph.GetDistance(startNode,endNode)

	return distance
}

func (graph *ItemGraph) addNodeIfDoesNotExist(value string) *Node {
	if !graph.nodeExist(value) {
		graph.AddNode(&Node{value: value})
	}
	return graph.getNode(value)
}

func (graph ItemGraph) getNode(s string) *Node {
	var node_found = &Node{value:""}
	for node := range graph.nodes{
		if graph.nodes[node].value == s {
			node_found = graph.nodes[node]
		}
	}
	return node_found
}

func (graph ItemGraph) nodeExist(s string) bool {
	for node := range graph.nodes{
		if graph.nodes[node].value == s {
			return true
		}
	}
	return false
}

func (graph *ItemGraph) canReach(EndNode *Node, StartNode *Node) bool {
	if StartNode.value == EndNode.value {
		return true
	}
	for _, nextNode := range graph.edges[*StartNode] {
			return graph.GetDistance(nextNode,EndNode) < NOT_FOUND
	}
	return false
}

func (graph *ItemGraph) GetDistance(StartNode *Node, EndNode *Node) int {

	nodeSeen := []*Node{}
	if StartNode == EndNode {
		return 0
	}
	return graph.GetDistanceRecursive(StartNode,EndNode,nodeSeen)
}

func (graph *ItemGraph) GetDistanceRecursive(StartNode *Node, EndNode *Node, NodeSeen []*Node) int {
	if StartNode == EndNode {
		return 0
	}
	results := []int{}
	for edge := range graph.edges[*StartNode] {
		nextNode := graph.edges[*StartNode][edge]
		if ListContainsNode(NodeSeen, nextNode){
			continue
		} else {
			NodeSeen = append(NodeSeen,nextNode)
		}
		results = append(results,graph.GetDistanceRecursive(nextNode,EndNode,NodeSeen))
	}
	for _, node_obj := range graph.nodes {
		for _, edge_obj := range graph.edges[*node_obj] {
			if edge_obj == StartNode {
				if ListContainsNode(NodeSeen, node_obj){
					continue
				} else {
					NodeSeen = append(NodeSeen,node_obj)
				}
				results = append(results,graph.GetDistanceRecursive(node_obj,EndNode,NodeSeen))
			}
		}
	}
	if len(results)>0{
		return Min(results) + 1
	}
	return NOT_FOUND
}

func ListContainsNode(list []*Node, node *Node) bool {
	for _, element := range list {
		if element == node {
			return true
		}
	}
	return false
}

func Min(multiple_int []int) int {
	smallest := multiple_int[0]
	for _, element :=range multiple_int{
		if smallest > element {
			smallest = element
		}
	}
	return smallest
}