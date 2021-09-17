package util

import "fmt"

func NNTable() {
	for j := 9; j > 0; j-- {
		for i := 1; i < j+1; i++ {
			fmt.Printf("%v*%v=%v;", j, i, j*i)
		}
		fmt.Println()
	}
}

func SwitchDemo() {
	a := 'a'
	n := 0
	switch a {
	case 'b':
		n++
		fmt.Println("b 停止了么")
	case 'a':
		n++
		fmt.Println("a 停止了么")
	case 'c':
		n++
		fmt.Println("c 停止了么")
	}
	fmt.Println("n:", n)
}

func scopeTest1() int {
	x := 5
	defer func() { x++ }()
	return x
}

func scopeTest2() (x int) {
	defer func() { x++ }()
	return 5
}

func scopeTest3() (y int) {
	x := 5
	defer func() { x++ }()
	return x
}

func scopeTest4() (x int) {
	defer func(x int) { x++ }(x)
	return 5
}

func ScopeTest() {
	fmt.Println(scopeTest1())
	fmt.Println(scopeTest2())
	fmt.Println(scopeTest3())
	fmt.Println(scopeTest4())
}

func DeferTestCalc(s string, a, b int) int {
	ret := a + b
	fmt.Println(ret, a, b)
	return ret
}
