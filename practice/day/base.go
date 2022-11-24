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

// D15 字符串无法修改，[]byte 可以修改
func D15() {
	str := "hello"

	a := []byte(str)
	a[0] = 'T'

	fmt.Println(string(a))
}

func D16() {
	// map 要用make分配空间
	//var m map[string]int
	var m = make(map[string]int)
	m["a"] = 1

	if v, ok := m["n"]; ok {
		fmt.Println(v)
	}
}

type a interface {
	showA()
}
type b interface {
	showB()
}
type worker struct {
	i int
}

func (w worker) showA() {
	fmt.Println(w.i + 10)
}

func (w worker) showB() {
	fmt.Println(w.i + 20)
}

// D17 虽然var声明的变量类型，但还是可以通过断言为其它更丰富的类型
func D17() {
	var a1 a = worker{3}
	s := a1.(worker)
	s.showA()
	s.showB()
}

// D18 defer执行顺序、值参数和指针参数
func D18() {
	type person struct {
		age int
	}
	p := &person{28}

	defer fmt.Println(p.age)

	defer func(p *person) {
		fmt.Println(p.age)
	}(p)

	defer func() {
		fmt.Println(p.age)
	}()

	p.age = 29
}

// D19 结构体是指针类型 重新赋值会导致变量指向的地址更改
func D19() {
	type person struct {
		age int
	}
	p := &person{28}

	defer fmt.Println(p.age)

	defer func(p *person) {
		fmt.Println(p.age)
	}(p)

	defer func() {
		fmt.Println(p.age)
	}()

	p = &person{29}
}

type person interface {
	speak(s string) string
}

type pp struct {
}

func (p *pp) speak(param string) (talk string) {
	return param
}

func D20() {
	// 因为pp实现的时指针类型的方法speak，
	// 所以赋值时只能赋值指针类型
	//var p person = pp{}
	var p person = &pp{}
	fmt.Println(p.speak("23423"))
}

// D21 当动态值和动态类型都为 nil 时，接口类型值才为 nil
func D21() {
	// 我猜测 指针类型未开辟空间 所以为nil
	var a *pp
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil", a)
	}
	fmt.Println(reflect.TypeOf(a), "a")

	// https://www.topgoer.cn/docs/gomianshiti/mian26 说是因为动态类型不是nil 但是通过反射打印类型 a,b是相同的
	var b person = a
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil", b)
	}
	fmt.Println(reflect.TypeOf(a)) // *day.pp
	fmt.Println(reflect.TypeOf(b)) // *day.pp
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

func (d direction) String() string {
	return [...]string{"north", "east", "south", "west"}[d]
}

// D22 iota String 方法
func D22() {
	fmt.Println(west)
}

func D23() {
	type math struct {
		x, y int
	}
	//m := map[string]math{"h": {1, 2}}
	// 改为指针类型
	m := map[string]*math{"h": &math{1, 2}}

	// 因为 map[string]math 是值类型，所以不能直接赋值
	m["h"].x = 3

	fmt.Println(m["h"])

	//// 值类型赋值，不会改变原值
	//m2 := m["h"]
	//m2.x = 3
	//
	//// 使用零时变量提黄 键值
	//m["h"] = m2
	//
	//fmt.Println(m)
}

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

// D24 变量作用域。
// 问题出在操作符:=，对于使用:=定义的变量，
// 如果新变量与同名已定义的变量不在同一个作用域中，
// 那么 Go 会新定义这个变量。
// 对于本例来说，main() 函数里的 p 是新定义的变量，
// 会遮住全局变量 p，导致执行到bar()时程序，全局变量 p 依然还是 nil，程序随即 Crash。
func D24() {
	var err error
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}

// D25
// 第一步执行r = n +1，
// 接着执行第二个 defer，由于此时 f() 未定义，引发异常，
// 随即执行第一个 defer，异常被 recover()，程序正常执行，
// 最后 return。
func D25() {
	var f = func(n int) (r int) {
		defer func() {
			r += n
			recover()
		}()
		var e func()

		defer e()
		e = func() {
			r += 2
		}
		return n + 1
	}
	fmt.Println(f(3))
}
