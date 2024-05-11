package graph

type Edge struct {
	ID         string
	Start      *Node
	End        *Node
	Properties map[string]interface{}
}
