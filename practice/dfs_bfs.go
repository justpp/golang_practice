package practice

import (
	"container/list"
	"fmt"
	"giao/pkg/util"
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
				newPath := make([]string, len(path))
				copy(newPath, path)
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
	human   int
}

func (s state) isValid() bool {

	// 狼羊同岸 且农夫不在
	if s.wolf == s.goat && s.goat != s.human {
		return false
	}

	if s.goat == s.cabbage && s.cabbage != s.human {
		return false
	}
	return true
}

func (s state) isFinal() bool {
	return s.human == 0 && util.ArrCompare([]int{
		s.wolf,
		s.goat,
		s.cabbage,
		s.human,
	})

}

func bfsWolfSheep(start state) []state {
	visited := make(map[state]bool)
	queue := list.New()
	queue.PushBack([]state{start})
	visited[start] = true

	times := 0

	for queue.Len() > 0 {
		curr := queue.Front().Value.([]state)
		queue.Remove(queue.Front())
		times++

		lastStep := curr[len(curr)-1]
		if lastStep.isFinal() {
			return curr
		}

		for _, next := range getNextStates(lastStep) {
			if !visited[next] && next.isValid() {
				visited[next] = true

				// 记录多种可能的路线
				curr = append(curr, next)
				queue.PushBack(curr)
			}
		}
	}
	return nil
}

func getNextStates(s state) []state {

	human := s.human
	if s.human == 1 {
		human = 0
	} else {
		human = 1
	}
	var next []state

	// human 往返
	next = append(next, state{s.wolf, s.goat, s.cabbage, human})

	// wolf crosses alone
	if s.human == s.wolf {
		if s.wolf == 1 {
			next = append(next, state{s.wolf - 1, s.goat, s.cabbage, human})
		} else {
			next = append(next, state{s.wolf + 1, s.goat, s.cabbage, human})
		}
	}

	// goat crosses alone
	if s.human == s.goat {
		if s.goat == 1 {
			next = append(next, state{s.wolf, s.goat - 1, s.cabbage, human})
		} else {
			next = append(next, state{s.wolf, s.goat + 1, s.cabbage, human})
		}
	}

	// cabbage crosses alone
	if s.human == s.cabbage {
		if s.cabbage == 1 {
			next = append(next, state{s.wolf, s.goat, s.cabbage - 1, human})
		} else {
			next = append(next, state{s.wolf, s.goat, s.cabbage + 1, human})
		}
	}

	fmt.Println("curr", s, "next", next)
	return next
}

func BfsWolfSheep() {
	start := state{1, 1, 1, 1}
	fmt.Println(bfsWolfSheep(start))
}
