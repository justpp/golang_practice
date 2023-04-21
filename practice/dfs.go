package practice

import "fmt"

// DfsAll dfs 全排 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func DfsAll(arr []int) {

	n := len(arr)
	book := make([]int, n)
	lineRes := make([]int, n)
	var loopDep func(step int)
	count := 0

	loopDep = func(step int) {
		if step == n {
			fmt.Println(lineRes)
			count++
			return
		}
		for i := 0; i < n; i++ {
			if book[i] != 1 {
				lineRes[step] = arr[i]
				book[i] = 1
				loopDep(step + 1)
				book[i] = 0
			}
		}
	}
	loopDep(0)
	fmt.Println("count:", count)
}
