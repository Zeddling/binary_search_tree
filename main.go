package main

import (
	"fmt"
	"io"
	"os"
)

//	Represents a node in the tree.
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

//	Stores the pointer to the root node of
//	the tree.
type BinaryTree struct {
	root *BinaryNode
}

//	Describes functions for all Node structs
type Node interface {
	/*
		Creates new Node and inserts data on tree.
		Inserts data less than root to the left
		and vice versa to the right
	*/
	New(int64)
}

//	Interface to describe all functions supported by
//  the tree
type Tree interface {
	/*
		Creates a tree and a root node
	*/
	Insert(int64) Tree
}

/*
	Asserts if both trees are identical

	Params:
		Root nodes for both tress
*/
func IsSameTree(root1 *BinaryNode, root2 *BinaryNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	} else if root1.data == root2.data &&
		IsSameTree(root1.left, root2.left) &&
		IsSameTree(root1.right, root2.right) {
		//	Traverses entire tree and returns true if
		//	this is the only condition that is satisfied
		return true
	}
	return false
}

/*
	Creates new Node and inserts data on tree.
	Inserts data less than root to the left
	and vice versa to the right.

	Params:
		data -> int64 number
*/
func (n *BinaryNode) New(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{
				data:  data,
				left:  nil,
				right: nil,
			}
		} else {
			n.left.New(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{
				data:  data,
				left:  nil,
				right: nil,
			}
		} else {
			//  Traverse tree till satisfied conditions
			//	for insertion are met
			n.right.New(data)
		}
	}
}

/*
	Deletes a node
		Case 1: Node is a leaf node
		Case 2: Node has one child
		Case 3: Node has both children

	Params:
		Node to be deleted
*/
func Delete(root *BinaryNode, key int64) *BinaryNode {
	if root == nil {
		return root
	} else if key < root.data {
		root.left = Delete(root.left, key)
	} else if key > root.data {
		root.right = Delete(root.right, key)
	} else {
		if root.left == nil && root.right == nil {
			//	Case 1
			root = nil
		} else if root.left == nil {
			//	Case 2
			root = root.right
		} else if root.right == nil {
			//	Case 2
			root = root.left
		} else {
			//	Case 3
			temp := root.right

			for {
				if temp.left == nil {
					break
				}
				temp = temp.left
			}

			root.data = temp.data
			root.right = Delete(root.right, temp.data)
		}
	}
	return root
}

/*
	Creates a tree and a root node
	Params:
		data -> int64 number
*/
func (t *BinaryTree) Insert(data int64) Tree {
	if t.root == nil {
		t.root = &BinaryNode{
			data:  data,
			left:  nil,
			right: nil,
		}
	} else {
		t.root.New(data)
	}
	return t
}

/*
	Prints tree from the provided node
		Params:
		1. w :any writer of type io.Writer
		2. space: Spacing degree for the pretty print.
		3. node: the node from which the tree will be printed from
		4. ch: a character for denoting which side of the current is being printed
*/
func Print(w io.Writer, space int, node *BinaryNode, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < space; i++ {
		fmt.Fprint(w, " ")
	}

	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	Print(w, space+2, node.left, 'L')
	Print(w, space+2, node.right, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree2 := &BinaryTree{}
	tree.Insert(100).
		Insert(-90).
		Insert(33).
		Insert(12).Insert(25)

	tree2.Insert(100).
		Insert(-90).
		Insert(33).
		Insert(12).Insert(25)

	Print(os.Stdout, 0, tree.root, 'M')
	Print(os.Stdout, 0, tree2.root, 'M')

	fmt.Println(
		"Both trees are identical: ",
		IsSameTree(tree.root, tree2.root),
	)

	Delete(tree.root, 12)
	fmt.Println("After deleting a key")
	Print(os.Stdout, 0, tree.root, 'M')
}
