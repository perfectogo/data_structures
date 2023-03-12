package ingolang

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func (l *List) Print() {
	var head *Node = l.Head

	fmt.Print("[")

	for head != nil {
		fmt.Print(head.Value, ", ")
		head = head.Next
	}

	fmt.Print("]")
}
