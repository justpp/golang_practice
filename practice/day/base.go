package day

import (
	"fmt"
	"reflect"
	"time"
)

// D1 延迟函数defer 先进后出 最后输出panic
func D1() {
	defer func() { fmt.Println("前") }()
	defer func() { fmt.Println("中") }()
	defer func() { fmt.Println("后") }()
	panic("panic err")
}

// D2 Two for rang val是一个单独的遍历，每轮引用变量的副本
func D2() {
	arr := []int{1, 2, 3, 4}
	m := make(map[int]*int)
	for i, v := range arr {
		// 使用v变量存储副本值
		// 这里&v是变量v的地址而不是arr[i]的地址
		m[i] = &v
		// 可以创建新变量赋值
		//val := v
		//m[i] = &val
	}
	for i, i2 := range m {
		fmt.Println(i, *i2)
	}
}

// D4 声明一次性结构体 *解引用赋值
func D4() {
	a := &struct {
		name string
	}{"234"}
	a.name = "000"
	fmt.Println(a)
	(*a).name = "999"
	fmt.Println(a)
}

type MyInt1 int
type MyInt2 = int

// D5 类型别名和新类型
func D5() {
	var i int = 0
	// cannot use i (type int) as type MyInt1 in assignment
	//var i1 MyInt1 = i
	// 可改写
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

// D6 接口类型可以赋值nil
func D6() {
	var i interface{} = nil
	fmt.Println(i)
}

// D7  类型转换
func D7() {
	var myInt MyInt1 = 4
	newInt := int(myInt)
	println(newInt)
}

// D8  chan 声明
// 若chan中没有数据，发生读阻塞
// 若chan已满，发生写阻塞
func D8() {
	var a chan int     // 只是为变量a声明类型，还需要为其make开辟空间才能试用
	a = make(chan int) // 若声明时未指定buff,必须有协程在读或协程在写，相当于快递不能放快递柜，只能送货上门
	go func() {
		time.Sleep(time.Second)
		a <- 2
		println("in")
	}()
	// 读阻塞，直到chan写入
	println("read", <-a)
}

// D9
// 访问map中不存在的键值，返回相应类型的零值
func D9() {
	type person struct {
		name string
	}
	var m map[person]int
	p := person{"234"}
	println(m[p])
}

// D10 可变长参数
func D10() {
	var a = func(args ...int) {
		fmt.Println(args[0])
	}
	var arr = []int{1, 2, 3, 4}
	a(arr...)
	println(arr[0])
}

// D11 不同类型不能相加
func D11() {
	a := 2
	b := 4.14
	fmt.Println(float64(a) + b)
}

// D12 interface类型的变量可以使用 i.()断言
func D12() {
	var i interface{}
	i = 2
	f, ok := i.(interface{})
	if ok {
		fmt.Println("ok", ok)
	}
	fmt.Println(f)
}

// D13 切片操作 [i:j:k] i起点(从0开始) j终点(默认为数组的长度，不包含当前位置) k容量
func D13() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr len", len(a))
	b := a[2:4:5]
	fmt.Println("b ", b)
}

// D14 %+输出数值的符号
func D14() {
	a := -5
	b := +5
	fmt.Printf("%d %+d b type:%s", a, b, reflect.ValueOf(b).Type())
}
