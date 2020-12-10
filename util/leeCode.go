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
	l := &ListNode{0, nil}
	tmpL := l
	addition := 0
	count := 0
	var a, b, val int
	fmt.Println("--------------------- ")
	fmt.Println("l1: ")
	ShowListNode(l1)
	fmt.Println("l2: ")
	ShowListNode(l2)
	fmt.Println("---------start------------ ")
	for (l1 != nil) || (l2 != nil) || addition > 0 {
		count++
		fmt.Println(l1, l2, addition)
		fmt.Println(count)
		if l1 != nil {
			a = l1.Val
		} else {
			a = 0
		}
		if l2 != nil {
			b = l2.Val
		} else {
			b = 0
		}
		val = a + b + addition
		fmt.Printf("a:%v b:%v addition:%v \n", a, b, addition)

		if val > 9 {
			val -= 10
			addition = 1
		} else {
			addition = 0
		}

		tmpL.Val = val
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if (l1 != nil) || (l2 != nil) || addition > 0 {
			tmpL.Next = &ListNode{0, nil}
			tmpL = tmpL.Next
		}

		fmt.Println("tmpL: ")
		ShowListNode(tmpL)
		fmt.Println("L1: ")

		ShowListNode(l1)

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
	for node != nil {
		showSlice = append(showSlice, node.Val)
		node = node.Next
	}
	fmt.Println(showSlice)
	return
}
