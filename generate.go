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
