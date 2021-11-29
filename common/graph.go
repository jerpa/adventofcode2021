package common

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Graph is a simple graph database
type Graph struct {
	nodes map[string]*GraphNode
}

// AddNode adds a node to a graph
func (g *Graph) AddNode(node GraphNode) (*GraphNode, error) {
	if g.nodes == nil {
		g.nodes = map[string]*GraphNode{}
	}
	if node.edgesFrom == nil {
		node.edgesFrom = map[string][]*GraphEdge{}
	}
	if node.edgesTo == nil {
		node.edgesTo = map[string][]*GraphEdge{}
	}
	if node.ID == "" {
		node.ID = uuid.New().String()
	}
	if _, exists := g.nodes[node.ID]; exists {
		return nil, errors.New("Node already exists")
	}
	g.nodes[node.ID] = &node

	return &node, nil
}

// GetNodeByID returns a node
func (g *Graph) GetNodeByID(id string) (*GraphNode, error) {
	if _, exists := g.nodes[id]; !exists {
		return nil, errors.New("Node does not exists")
	}
	return g.nodes[id], nil
}

// GetNodesByType returns a list of nodes
func (g *Graph) GetNodesByType(t string) []*GraphNode {
	result := []*GraphNode{}
	for k := range g.nodes {
		if g.nodes[k].Type == t {
			result = append(result, g.nodes[k])
		}
	}
	return result
}
func (g *Graph) RemoveNode(node *GraphNode) error {
	if _, ok := g.nodes[node.ID]; !ok {
		return errors.New("Unknown node")
	}
	delete(g.nodes, node.ID)
	for _, v := range g.nodes {
		for k := range v.edgesTo {
			for _, e := range v.edgesTo[k] {
				if e.To == node {
					v.RemoveEdgeTo(k, *e)
				}
			}
		}
		for k := range v.edgesFrom {
			for _, e := range v.edgesFrom[k] {
				if e.From == node {
					v.RemoveEdgeFrom(k, *e)
				}
			}
		}
	}
	return nil
}

// Search("id=apa,")
func (g *Graph) Search(q string) []*GraphNode {
	result := []*GraphNode{}
	nodes := []*GraphNode{}
	qd := strings.Split(q, ",")
	qp := strings.Split(qd[0], "=")
	if qp[0] == "id" {
		k, err := g.GetNodeByID(qp[1])
		if err != nil {
			panic(err.Error())
		}
		nodes = append(nodes, k)
	} else if qp[0] == "type" {
		k := g.GetNodesByType(qp[1])

		nodes = append(nodes, k...)
	} else {

		//g.GetNodesByType()

	}
	return result
}

// GraphEdge is an edge in a Graph
type GraphEdge struct {
	ID       string
	From     *GraphNode
	FromName string
	To       *GraphNode
}

// GraphNode is a node in a Graph
type GraphNode struct {
	ID        string
	Type      string
	edgesTo   map[string][]*GraphEdge
	edgesFrom map[string][]*GraphEdge
}

func (gn *GraphNode) AddEdge(name string, node *GraphNode) error {
	id := uuid.New().String()

	ge := GraphEdge{ID: id, From: gn, To: node, FromName: name}

	if _, ok := gn.edgesTo[name]; !ok {
		gn.edgesTo[name] = []*GraphEdge{}
	}
	gn.edgesTo[name] = append(gn.edgesTo[name], &ge)

	if _, ok := node.edgesFrom[name]; !ok {
		node.edgesFrom[name] = []*GraphEdge{}
	}
	node.edgesFrom[name] = append(node.edgesFrom[name], &ge)

	return nil
}
func (gn *GraphNode) RemoveEdgeTo(name string, edge GraphEdge) error {
	if _, ok := gn.edgesTo[name]; !ok {
		return errors.New("Invalid name")
	}
	for i := len(gn.edgesTo[name]) - 1; i >= 0; i-- {
		if gn.edgesTo[name][i].ID == edge.ID {
			n := gn.edgesTo[name][i].To
			gn.edgesTo[name] = append(gn.edgesTo[name][:i], gn.edgesTo[name][i+1:]...)
			n.RemoveEdgeFrom(name, edge)
			return nil
		}
	}
	return errors.New("Edge not found")
}
func (gn *GraphNode) RemoveEdgeFrom(name string, edge GraphEdge) error {
	if _, ok := gn.edgesFrom[name]; !ok {
		return errors.New("Invalid name")
	}
	for i := len(gn.edgesFrom[name]) - 1; i >= 0; i-- {
		if gn.edgesFrom[name][i].ID == edge.ID {
			n := gn.edgesFrom[name][i].From
			gn.edgesFrom[name] = append(gn.edgesFrom[name][:i], gn.edgesFrom[name][i+1:]...)
			n.RemoveEdgeTo(name, edge)
			return nil
		}
	}
	return errors.New("Edge not found")
}
func (gn *GraphNode) ConnectedTo(name string, node *GraphNode) bool {
	for _, v := range gn.edgesTo[name] {
		if v.To == node {
			return true
		}
	}
	return false
}
func (gn *GraphNode) GetEdgesTo(name string) []*GraphEdge {
	return gn.edgesTo[name]
}
func (gn *GraphNode) GetEdgesFrom(name string) []*GraphEdge {
	return gn.edgesFrom[name]
}
