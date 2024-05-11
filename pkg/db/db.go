package db

import (
	"graphdb-mvp/pkg/graph"
)

type GraphDB struct {
	Nodes map[string]*graph.Node
	Edges map[string]*graph.Edge
}

func NewGraphDB() *GraphDB {
	return &GraphDB{
		Nodes: make(map[string]*graph.Node),
		Edges: make(map[string]*graph.Edge),
	}
}

func (db *GraphDB) AddNode(node *graph.Node) {
	db.Nodes[node.ID] = node
}

func (db *GraphDB) AddEdge(edge *graph.Edge) {
	db.Edges[edge.ID] = edge
	edge.Start = db.Nodes[edge.Start.ID]
	edge.End = db.Nodes[edge.End.ID]
}
