package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value rune
	Right *Tree
}

type Queue []*Tree

func (q *Queue) Push(n *Tree) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n *Tree) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	return len(*q)
}

type Stack []*Tree

func (q *Stack) Push(n *Tree) {
	*q = append(*q, n)
}

func (q *Stack) Pop() (n *Tree) {
	x := q.Len() - 1
	n = (*q)[x]
	*q = (*q)[:x]
	return
}
func (q *Stack) Len() int {
	return len(*q)
}

func main() {
	exp := "123+-"
	fmt.Println("exp is: " + exp)
	myTree := postfixToTree(exp)
	fmt.Println("Tree: " + myTree.String())
	fmt.Println("Pre-Order")
	preOrderTraversal(myTree)
	fmt.Printf("\n")
	fmt.Println("In-Order")
	inOrderTraversal(myTree)
	fmt.Printf("\n")
	fmt.Println("Post-Order")
	postOrderTraversal(myTree)
	fmt.Printf("\n")

}

func postfixToTree(exp string) *Tree {
	myStack := &Stack{}
	for _, element := range exp {
		if element > 47 && element < 58 { //It is a number
			leafNode := &Tree{nil, element, nil}
			myStack.Push(leafNode)
		}
		if element == 47 || element == 42 || element == 43 || element == 45 { //It is * - / or +
			rightNode := myStack.Pop()
			leftNode := myStack.Pop()
			rootNode := &Tree{leftNode, element, rightNode}
			myStack.Push(rootNode)
		}
	}
	finalTree := myStack.Pop()
	return finalTree
}

func preOrderTraversal(t *Tree) {
	if t != nil {
		fmt.Printf("%c ", t.Value)
		preOrderTraversal(t.Left)
		preOrderTraversal(t.Right)
	}
}

func inOrderTraversal(t *Tree) {
	if t != nil {
		inOrderTraversal(t.Left)
		fmt.Printf("%c ", t.Value)
		inOrderTraversal(t.Right)
	}
}

func postOrderTraversal(t *Tree) {
	if t != nil {
		postOrderTraversal(t.Left)
		postOrderTraversal(t.Right)
		fmt.Printf("%c ", t.Value)
	}
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprintf("%c", t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
