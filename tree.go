// Package tree contains utils to create an manipulate a Tree
package tree

// Item contains relevant information about a file, or folder
type Item struct {
	name, path string
}

// ItemData is an interface for assessing name and path of objects representing
// files or folders
type ItemData interface {
	GetName() string
	GetPath() string
}

// Node defines a Node-like structure
type Node struct {
	item     Item
	children []*Node
	parent   *Node
}

// GetName returns the name of the input item
func GetName(item Item) string {
	return item.name
}

// GetPath returns the path of the input item
func GetPath(item Item) string {
	return item.path
}
