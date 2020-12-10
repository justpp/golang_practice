package util

import "fmt"

// 二维切片中两数之和 返回下标
func TwoSumReturnK(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

// 两数之和 两个倒叙链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func TowSumList(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	tmpL := l
	L1 := l1
	L2 := l2
	addition := 0
	count := 0
	fmt.Println("--------------------- ")
	fmt.Println("l1: ")
	ShowListNode(l1)
	fmt.Println("l2: ")
	ShowListNode(l2)
	fmt.Println("---------start------------ ")
	for (L1 != nil && L1.Next != nil) || (L2 != nil && L2.Next != nil) || addition > 0 {
		count++
		fmt.Println(L1, L2, addition)
		fmt.Println(count)
		a := 0
		b := 0
		if L1 != nil {
			a = L1.Val
		}
		if L2 != nil {
			b = L2.Val
		}
		val := a + b + addition
		fmt.Printf("a:%v b:%v addition:%v \n", a, b, addition)

		if val > 9 {
			val -= 10
			addition = 1
		} else {
			addition = 0
		}

		tmpL.Val = val
		if L1 != nil {
			L1 = L1.Next
		}
		if L2 != nil {
			L2 = L2.Next
		}
		tmpL.Next = new(ListNode)
		tmpL = tmpL.Next

		fmt.Println("tmpL: ")
		ShowListNode(tmpL)
		fmt.Println("L1: ")

		ShowListNode(L1)

		fmt.Println("l: ")
		ShowListNode(l)
		fmt.Println("------------")
	}
	return l
}

func ShowListNode(node *ListNode) {
	if node == nil {
		return
	}
	var showSlice = make([]int, 0)
	for node.Next != nil {
		showSlice = append(showSlice, node.Val)
		node = node.Next
	}
	fmt.Println(showSlice)
	return
}
