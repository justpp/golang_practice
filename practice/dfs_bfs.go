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

type state struct {
	wolf    int
	goat    int
	cabbage int
}

func (s state) isValid() bool {
	if s.wolf == s.goat && s.goat != s.cabbage {
		return false
	}
	if s.goat == s.cabbage && s.goat != s.wolf {
		return false
	}
	return true
}

func (s state) isFinal() bool {
	return s.wolf == 0 && s.goat == 0 && s.cabbage == 0
}

func bfsWolfSheep(start state) bool {
	visited := make(map[state]bool)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		curr := queue.Front().Value.(state)
		queue.Remove(queue.Front())

		fmt.Println(curr)
		if curr.isFinal() {
			return true
		}

		for _, next := range getNextStates(curr) {
			if !visited[next] {
				visited[next] = true
				queue.PushBack(next)
			}
		}
	}

	return false
}

func getNextStates(s state) []state {
	var next []state

	// wolf crosses alone
	if s.wolf == 1 {
		next = append(next, state{s.wolf - 1, s.goat, s.cabbage})
	} else {
		next = append(next, state{s.wolf + 1, s.goat, s.cabbage})
	}

	// goat crosses alone
	if s.goat == 1 {
		next = append(next, state{s.wolf, s.goat - 1, s.cabbage})
	} else {
		next = append(next, state{s.wolf, s.goat + 1, s.cabbage})
	}

	// cabbage crosses alone
	if s.cabbage == 1 {
		next = append(next, state{s.wolf, s.goat, s.cabbage - 1})
	} else {
		next = append(next, state{s.wolf, s.goat, s.cabbage + 1})
	}

	return next
}

func BfsWolfSheep() {
	start := state{1, 1, 1}
	fmt.Println(bfsWolfSheep(start))
}
