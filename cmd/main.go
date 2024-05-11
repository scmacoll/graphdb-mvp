// File path: cmd/main.go

package main

import (
	"fmt"
	"graphdb-mvp/pkg/db"
	"graphdb-mvp/pkg/graph"
)

func main() {
	gdb := db.NewGraphDB()

	// Create nodes
	node1 := &graph.Node{ID: "1", Properties: map[string]interface{}{"name": "Alice"}}
	node2 := &graph.Node{ID: "2", Properties: map[string]interface{}{"name": "Bob"}}

	// Add nodes to the database
	gdb.AddNode(node1)
	gdb.AddNode(node2)

	// Create and add an edge
	edge := &graph.Edge{ID: "1-2", Start: node1, End: node2, Properties: map[string]interface{}{"relation": "friends"}}
	gdb.AddEdge(edge)

	fmt.Println("Graph Database Nodes:", gdb.Nodes)
	fmt.Println("Graph Database Edges:", gdb.Edges)
}
