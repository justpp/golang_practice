package util

import "fmt"

var (
	// Coins 金币
	Coins = 100
	// Users ren
	users = []string{"MHatthew", "sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "ElizabetIh"}
	// Distribution gan
	Distribution = make(map[string]int, len(users))
)

// Ref 递归
func Ref(a int) int {
	if a == 1 {
		return 1
	}
	return a * Ref(a-1)
}

// Dispatchicons 分金币
func Dispatchicons() int {
	for _, name := range users {
		for _, s := range name {
			switch s {
			case 'e', 'E':
				Distribution[name]++
			case 'i', 'I':
				Distribution[name] += 2
			case 'o', 'O':
				Distribution[name] += 3
			case 'u', 'U':
				Distribution[name] += 4
			}
		}
		Coins -= Distribution[name]
	}

	return Coins
}

// GoStep 走台阶 一次走一步、一次走两步有多少种走法
func GoStep(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return GoStep(n-1) + GoStep(n-2)
}

// 自定义类型和类型别名

// MyInt giao 自定义类型
type MyInt int

// YourInt 类型别名 编写的过程中生效 执行之后没了 赋值只能在函数体内 单独写外面会报错 non-declaration
// tpye YourInt = int

// 结构体

// Human 结构体
type Human struct {
	Name  string
	Age   int
	Hobby []string
}

// 闭包
func UseFunc(f func()) {
	f()
}

func F1(i, j int) {
	fmt.Println(i)
	fmt.Println(j)
}
func Closure(int1, int2 int) func() {
	return func() {
		F1(int1, int2)
	}
}
