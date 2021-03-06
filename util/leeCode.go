package util

import (
	"fmt"
	"strings"
)

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

func TwoSumListTest() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	res := TowSumList(l1, l2)
	ShowListNode(res)
}

// 最长子串
func LongestCommSub(str1 string, str2 string) string {
	var data = make(map[int]map[int]int)
	var maxLength = 0
	var maxStr1 = 0
	var maxStr2 = 0

	for k1, v1 := range str1 {
		for k2, v2 := range str2 {
			// 二维map麻烦的很
			if _, ok := data[k1]; !ok {
				data[k1] = make(map[int]int)
			}
			if v1 == v2 {
				// 检查上一个相同的字是否存在
				if n, ok := data[k1-3][k2-3]; ok && n > 0 {
					data[k1][k2] = 1 + data[k1-3][k2-3] // 相同第二次 再上一次基础上加一
				} else {
					data[k1][k2] = 1 // 相同第一次 加一
				}
			} else {
				data[k1][k2] = 0 // 不相同 节点值0
			}
			if maxLength < data[k1][k2] {
				maxLength = data[k1][k2]
				maxStr1 = k1
				maxStr2 = k2
			}
		}
	}
	// 输出最大字串
	strSlice := make([]string, maxLength) // 搞一个maxLength容量的切片来存字串
	str1Slice := make(map[int]rune)
	for k, v := range str1 {
		str1Slice[k] = v
	}
	for true {
		if n, ok := data[maxStr1][maxStr2]; !ok || n == 0 {
			break
		}
		// 切片没有append扩容的情况下最大子键 maxLength-1
		strSlice[maxLength-1] = string(str1Slice[maxStr1])
		maxStr1 -= 3
		maxStr2 -= 3
		maxLength-- // 倒续存入
	}
	// 用strings包给他拼接起来
	subStr := strings.Join(strSlice, "")
	fmt.Println("result:", subStr)
	return subStr
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var L1 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	},
}

var L2 = &ListNode{
	Val: 5,
	Next: &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}

// 两个倒序链表相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res
	addition := 0
	for l1 != nil || l2 != nil || addition != 0 {
		curr.Next = &ListNode{}
		curr = curr.Next
		if l1 != nil {
			curr.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curr.Val += l2.Val
			l2 = l2.Next
		}
		curr.Val += addition
		addition = curr.Val / 10
		curr.Val %= 10
	}
	return res.Next
}

// 字符串内最长重复字串 滑窗
func LengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

//  查找两个正序数组的中位数  给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sli := append(nums1, nums2...)
	l := len(sli)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
	if l%2 == 0 {
		return float64(sli[l/2]+sli[l/2-1]) / 2
	}
	return float64(sli[l/2])
}
