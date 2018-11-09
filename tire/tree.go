package tire

import (
	"net/http"
	"strings"
)

const PathSplit = "/"

type Tree struct {
	Root     *Node
	basePath string
	Depth    int
}

func NewTree(basePath string) *Tree {
	return &Tree{
		basePath: basePath,
	}
}

func (t *Tree) AddNode(path string, method string, handle http.HandlerFunc) {
	pattern := strings.Split(path, PathSplit)
	if len(pattern) == 0 {
		//add for local
		if t.Root == nil {
			t.Root = NewNode(t.basePath, t.Depth)
		}
		if t.Root.IsMethodNotExist(method) {
			t.Root.AddHandle(method, handle)
		}

	} else {
		//create new node and add depth,count
		for index, value := range pattern {
			if val := t.Root.Sub[value]; val != nil {
				if val.IsMethodNotExist(method) {

				} else {

				}
				val.Count = val.Count + 1
			} else {
				newNode := NewNode(value, t.Depth+index+1)
				newNode.AddHandle(method, handle)
			}
		}
	}
}

func (t *Tree) FindNode(method string, path string) *Node {
	pattern := strings.Split(path, "/")
	currentNode := t.Root
	for _, currentpattern := range pattern {
		currentNode.MatchHandle(method, currentpattern)
	}
	return nil
}

type Node struct {
	Pattern string
	Count   int
	Depth   int
	Sub     map[string]*Node
	handle  []http.HandlerFunc
}

func (n *Node) IsMethodNotExist(method string) bool {
	switch strings.ToUpper(method) {
	case http.MethodGet:
		return n.handle[0] == nil
	case http.MethodHead:
		return n.handle[1] == nil
	case http.MethodPost:
		return n.handle[2] == nil
	case http.MethodPut:
		return n.handle[3] == nil
	case http.MethodDelete:
		return n.handle[4] == nil
	case http.MethodOptions:
		return n.handle[5] == nil
	}
	return false
}

func (n *Node) AddHandle(method string, handle http.HandlerFunc) {
	switch strings.ToUpper(method) {
	case http.MethodGet:
		n.handle[0] = handle
	case http.MethodHead:
		n.handle[1] = handle
	case http.MethodPost:
		n.handle[2] = handle
	case http.MethodPut:
		n.handle[3] = handle
	case http.MethodDelete:
		n.handle[4] = handle
	case http.MethodOptions:
		n.handle[5] = handle
	}
	return
}

func (n *Node) MatchHandle(method string, path string) bool {
	return true
}

func NewNode(pattern string, depth int) *Node {
	return &Node{
		Pattern: pattern,
		Count:   0,
		Depth:   depth,
		Sub:     make(map[string]*Node),
		handle:  make([]http.HandlerFunc, 6, 6),
	}
}
