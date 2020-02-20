# graph-test-data

Ready-to-use graph/network files in CSV for testing graph applications and algorithms.

You can also tweak the generating script to generate more graphs with your specified parameters (number of nodes etc.)

Feel free to make a PR to add your generated test graph.

## Available Graphs

Undirected Path (10, 1000, 10000 nodes)

Undirected Complete Graph (10, 100, 1000 nodes)

Undirected Grid (5\*5, 10\*10, 30*30 nodes)

Undirected Ladder (10, 100, 1000 nodes)
- ladders are just special cases of grids

(Make a PR to add more...)

## Run the Golang Graph File Generator

```
go run generate.go
```

CSV files will be generated under the `gen/` directory.