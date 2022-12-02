package day

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
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

func change(s ...int) {
	fmt.Printf("change before ptr %p\n", s)
	s = append(s, 3)
	fmt.Println("change", s)
	fmt.Printf("change after ptr %p\n", s)
}

// D26 可变长参数 接收值时slice
// 可变函数、append()操作。
// Go 提供的语法糖…，
// 可以将 slice 传进可变函数，不会创建新的切片。
// 第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变；
// 第二次调用change() 函数时，使用了操作符[i,j]获得一个新的切片，假定为 slice1，它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。
func D26() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	fmt.Printf("slice before ptr %p\n", slice)
	change(slice...)
	fmt.Printf("slice after ptr %p\n", slice)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}

// D27 forr 切片时拷贝的是指针的副本，其指向的还是相同的数组
func D27() {
	var a = []int{1, 2, 3, 4, 5}
	var r [6]int
	for i, i2 := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
			// append生成了新的底层arr
			a = append(a, 6)
			a[1] = 2
		}
		r[i] = i2
	}
	fmt.Println("a", a)
	fmt.Println("r", r)
}

// D28 map遍历顺序不确定性
func D28() {
	var m = map[string]int{
		"AA": 2,
		"BB": 3,
		"CC": 4,
	}
	count := 0
	for k, v := range m {
		if count == 0 {
			delete(m, "AA")
		}
		count++
		fmt.Println(k, v)
	}
	fmt.Println("count", count)
	var b bytes.Buffer
	b.WriteString("234")
	b.WriteString("666")
	fmt.Println(b.String())
}

// D30 强制类型转换
func D30() {
	type myInt int
	var i int = 4

	// 类型名称不同，不能直接赋值
	//var m myInt = i

	// 类型断言不是这个写法啊
	//var m myInt = (myInt)i

	// 类型断言 判断的是动态类型，不是类型转换啊
	//var m myInt = i.(int)

	// 强制类型转换
	var m myInt = myInt(i)
	fmt.Println(i, m)
}

// D31 switch
func D31() {
	a := 2
	switch a {
	case 1:
		fmt.Println("1")
	case 2, 4, 6, 7, 8: // 一个分支可以有多个case值
		fmt.Println("2")
		fallthrough // 通过fallthrough进入下一个case
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}

type integer1 int

func (i *integer1) Add(integer12 integer1) integer1 {
	return *i + integer12
}

// D32 类型断言 方法集
func D32() {
	var a integer1 = 2
	var b integer1 = 2

	var i interface{} = &a

	fmt.Println(i.(*integer1).Add(b))
}

func D33() {
	i := 1
	i++

	// 自增是语句，不是表达式，不能赋值给其他变量
	//j = i++

	// 没有i--
	// i--
	fmt.Println(i)
}

func D34() {
	// 声明切片时
	// make必须携带长度
	// 简单切片表达式 []int{...}必须有值
	// 通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	a := make([]int, 0)
	b := []int{1}
	var c []int
	c = append(c, 1)
	fmt.Println(a, b, c)
}

// D35 select 会随机选择一个可用通道
func D35() {
	runtime.GOMAXPROCS(1)
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch1 <- 1
	ch2 <- 2
	ch1 <- 5
	ch2 <- 5
	select {
	case value := <-ch1:
		fmt.Println("|", value)
	case <-ch2:
		//panic("2332")
	}
	if v, ok := <-ch1; ok {
		fmt.Println("ch1", v)
	}
	if v, ok := <-ch2; ok {
		fmt.Println("ch2", v)
	}
}

// D36
// 从一个已经关闭的chan中读取只能读到该类型的零值，
// 无论其有没有缓冲，都不会死锁
func D36() {
	ch := make(chan int)
	close(ch)
	fmt.Println(<-ch)
}

// D37 常量。
// 常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，所以常量无法寻址。
func D37() {
	const i = 9
	var j = 10
	//fmt.Println(&i)
	fmt.Println(&j)
}

type user struct{}
type user1 user
type user2 = user

func (u user1) m1() {
	fmt.Println("m1")
}

func (u user) m2() {
	fmt.Println("m2")
}

// D38
// type user1 user 创建了新类型
// type user2 = user 类型别名
// user1 并没有继承user的方法
func D38() {
	var u1 user1
	var u2 user2
	u1.m1()
	u2.m2()
}

// D39 interface 的内部结构，
// 我们知道接口除了有静态类型，还有动态类型和动态值，
// 当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。
// 这里的 x 的动态类型是 *int，所以 x 不为 nil。
func D39() {
	var f = func(x interface{}) {
		if x == nil {
			fmt.Println("empty")
			return
		}
		fmt.Println("not empty")
	}
	var i *int = nil
	f(i)
}

// D40 先定义下，第一个协程为 A 协程，第二个协程为 B 协程；
// 当 A 协程还没起时，主协程已经将 channel 关闭了，
// 当 A 协程往关闭的 channel 发送数据时会 panic，
// panic: send on closed channel。
func D40() {
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("in", i)
			ch <- i
		}
	}()

	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println("out", v)
			}
		}
	}()
	//close(ch)
	fmt.Println("end")
	time.Sleep(time.Second * 10)
}

// D41 有方向的chan不能关闭
func D41() {
	var f = func(ch <-chan int) {
		//close(ch)
	}
	ch := make(chan int)
	f(ch)
}

// D42 ，字面量初始化切片时候可以指定索引，
// 没有指定索引的元素会在前一个索引基础之上加一，
//所以输出[1 0 2 3]，而不是[1 3 2]
func D42() {
	var s = []int{2: 2, 3, 1: 1}
	fmt.Println(s)
}
