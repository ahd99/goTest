package main

import "fmt"

func main() {
	test1()
}

func test1() {
	list := createList(5)
	printList(list)
	fmt.Println()

	list = createList1([]int{10, 20, 30, 40, 60})
	printList(list)
	fmt.Println()

	list = createList1([]int{10, 20, 30, 40, 60})
	printList(list)
	fmt.Println()
}

func createList(a int) *Node {
	var pnode *Node
	for i := 0; i < a; i++ {
		node := &Node{i, pnode}
		pnode = node
	}
	return pnode
}

func createList1(a []int) *Node {
	cnode := new(Node)
	root := cnode
	var pnode *Node

	for _, v := range a {
		cnode.Value = v
		nnode := new(Node)
		cnode.Next = nnode
		pnode = cnode
		cnode = nnode
	}
	pnode.Next = nil
	return root
}

func createList2(a []int) *Node {
	cnode := &Node{Value: a[0], Next: nil}
	root := cnode

	for i := 1; i < len(a); i++ {
		nnode := new(Node)
		nnode.Value = a[i]
		cnode.Next = nnode
		cnode = nnode
	}
	return root
}

func printList(list *Node) {
	for list != nil {
		fmt.Printf("%d\t", list.Value)
		list = list.Next
	}
}

type Node struct {
	Value int
	Next  *Node
}
