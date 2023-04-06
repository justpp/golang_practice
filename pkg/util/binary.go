package util

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n Node) IsEmpty() bool {
	return n == Node{}
}

func (n *Node) Add(val int) {
	if n.IsEmpty() {
		n.Val = val
		return
	}
	if n.Val > val {
		if n.Right == nil {
			n.Right = &Node{}
		}
		n.Right.Add(val)
	} else {
		if n.Left == nil {
			n.Left = &Node{}
		}
		n.Left.Add(val)
	}
	return
}

func ShowBinary(node *Node) {
	PreorderTraversal(node)
}

// PreorderTraversal 前序遍历
func PreorderTraversal(node *Node) {
	if node == nil {
		return
	}
	if node.Left != nil {
		PreorderTraversal(node.Left)
	}
	fmt.Println(node.Val)
	if node.Right != new(Node) {
		PreorderTraversal(node.Right)
	}
	return
}

func (n *Node) search(val int) *Node  {
	if n.IsEmpty() {
		return nil
	}
	if n.Val == val {
		return n
	}
	if n.Val > val {
		if n.Right == nil {
			return nil
		}
		return n.Right.search(val)
	} else {
		if n.Left == nil {
			return nil
		}
		return n.Left.search(val)
	}
}

func (n *Node) MinDel() *Node  {
	if n.IsEmpty() {
		return nil
	}
	if n.Left == nil {
		return n.Right
	}
	return n
}

func CreateBinary() {
	var a = new(Node)
	arr := [...]int{9, 9, 10, 1, 3,2, 4, 5, 6}
	for _, i := range arr {
		a.Add(i)
	}
	ShowBinary(a)

}
