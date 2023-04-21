package practice

import (
	"container/list"
	"fmt"
)

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

func BFS() {
	graph := map[string][]string{
		"start": {"a", "b"},
		"a":     {"c", "d"},
		"b":     {"e", "f"},
		"c":     {"end"},
		"d":     {"end"},
		"e":     {"end"},
		"f":     {"end"},
	}
	fmt.Println(bfs(graph, "start", "end"))

}
func bfs(graph map[string][]string, start, end string) []string {
	l := list.New()
	l.PushBack([]string{start})
	visited := make(map[string]struct{})

	for l.Len() > 0 {
		path := l.Front().Value.([]string)
		l.Remove(l.Front())
		node := path[len(path)-1]
		if node == end {
			return path
		}
		for _, s := range graph[node] {
			if _, ok := visited[s]; !ok {
				visited[s] = struct{}{}
				var newPath []string
				newPath = append(newPath, path...)
				newPath = append(newPath, s)
				l.PushBack(newPath)
			}
		}
	}
	return nil
}
