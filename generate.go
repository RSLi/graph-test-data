package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

func main() {
	fmt.Println("Generating graph CSV files")

	EdgesToCSV(generateUndirectedPathGraph(10).Edges(), "gen/undirectedPath10.csv")
	EdgesToCSV(generateUndirectedPathGraph(1000).Edges(), "gen/undirectedPath1000.csv")
	EdgesToCSV(generateUndirectedPathGraph(10000).Edges(), "gen/undirectedPath10000.csv")

	EdgesToCSV(generateUndirectedGridGraph(5, 5).Edges(), "gen/undirectedGrid5x5.csv")
	EdgesToCSV(generateUndirectedGridGraph(10, 10).Edges(), "gen/undirectedGrid10x10.csv")
	EdgesToCSV(generateUndirectedGridGraph(30, 30).Edges(), "gen/undirectedGrid30x30.csv")

	EdgesToCSV(generateUndirectedLadderGraph(10).Edges(), "gen/undirectedLadder10.csv")
	EdgesToCSV(generateUndirectedLadderGraph(100).Edges(), "gen/undirectedLadder100.csv")
	EdgesToCSV(generateUndirectedLadderGraph(1000).Edges(), "gen/undirectedLadder1000.csv")

	EdgesToCSV(generateUndirectedCompleteGraph(10).Edges(), "gen/undirectedComplete10.csv")
	EdgesToCSV(generateUndirectedCompleteGraph(100).Edges(), "gen/undirectedComplete100.csv")
	EdgesToCSV(generateUndirectedCompleteGraph(1000).Edges(), "gen/undirectedComplete1000.csv")
}

// EdgesToCSV converts Gonum Graph edges into a CSV file.
func EdgesToCSV(edges graph.Edges, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"from", "to"})
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}

	for edges.Next() {
		edge := edges.Edge()
		fromNodeID := strconv.FormatInt(edge.From().ID(), 10)
		toNodeID := strconv.FormatInt(edge.To().ID(), 10)
		err := writer.Write([]string{fromNodeID, toNodeID})
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}

func generateUndirectedPathGraph(numberOfNodes int) *simple.UndirectedGraph {
	undirectedGraph := simple.NewUndirectedGraph()
	prev := undirectedGraph.NewNode()
	undirectedGraph.AddNode(prev)
	for i := 1; i < numberOfNodes; i++ {
		curr := undirectedGraph.NewNode()
		undirectedGraph.AddNode(curr)
		undirectedGraph.SetEdge(undirectedGraph.NewEdge(prev, curr))
		prev = curr
	}
	return undirectedGraph
}

func generateUndirectedGridGraph(r, c int) *simple.UndirectedGraph {
	undirectedGraph := simple.NewUndirectedGraph()
	list := []graph.Node{}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			newNode := undirectedGraph.NewNode()
			undirectedGraph.AddNode(newNode)
			if j != 0 {
				undirectedGraph.SetEdge(undirectedGraph.NewEdge(newNode, list[len(list)-1]))
			}
			if i != 0 {
				undirectedGraph.SetEdge(undirectedGraph.NewEdge(newNode, list[len(list)-c]))
			}
			list = append(list, newNode)
		}
	}
	return undirectedGraph
}

func generateUndirectedLadderGraph(numberOfRows int) *simple.UndirectedGraph {
	return generateUndirectedGridGraph(numberOfRows, 2)
}

func generateUndirectedCompleteGraph(numberOfNodes int) *simple.UndirectedGraph {
	undirectedGraph := simple.NewUndirectedGraph()
	list := []graph.Node{}
	for i := 0; i < numberOfNodes; i++ {
		newNode := undirectedGraph.NewNode()
		undirectedGraph.AddNode(newNode)
		for j := 0; j < len(list); j++ {
			undirectedGraph.SetEdge(undirectedGraph.NewEdge(newNode, list[j]))
		}
		list = append(list, newNode)
	}
	return undirectedGraph
}
