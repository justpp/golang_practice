package util

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"giao/calc"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// NNTable 九九乘法表
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

// ClosurePractice 闭包函数作用域练习
func ClosurePractice() {
	x := 1
	y := 2
	defer DeferTestCalc("AA", x, DeferTestCalc("A", x, y))
	x = 10
	defer DeferTestCalc("BB", x, DeferTestCalc("B", x, y))
	y = 20
	// A  3 1 2
	// B  12 10 2
	// BB 22 10 12
	// AA 4 1 3
}

func contextTest() {
	ctx, cancel := context.WithCancel(context.Background())
	go firstCtx(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("stop all sub goroutine")
	cancel()
	time.Sleep(5 * time.Second)
}

func firstCtx(ctx context.Context) {
	go secondCtx(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("first done")
			return
		default:
			fmt.Println("first running")
			time.Sleep(2 * time.Second)
		}
	}
}

func secondCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("second done")
			return
		default:
			fmt.Println("second running")
			time.Sleep(2 * time.Second)
		}
	}
}

type student struct {
	name string
	age  int
}

// StructPractice1 结构体练习
func StructPractice1() {
	m := make(map[string]*student)
	stubs := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	for _, stu := range stubs {
		s := stu // 若没有新创建变量储存 结果将为切片中最后一个值
		// 这是因为for range 中以同一片内存接收变量
		// 小王子 => *大王八
		// 娜扎 => *大王八
		// 大王八 => *大王八
		m[stu.name] = &s
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

type class struct {
	Title   string
	Student []*Stud
}

//Stud 学生
type Stud struct {
	ID     int `json:"id"`
	Gender string
	Name   string
}

// StructJsonPractice 结构体转json练习
func StructJsonPractice() {
	c := &class{
		"那一个班级",
		make([]*Stud, 0, 10),
	}
	for i := 0; i < 10; i++ {
		stu := &Stud{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Student = append(c.Student, stu)
	}
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("json:%s\n", data)

	str := `{"Title":"那一个班级","Student":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("un err:", err)
	}
	fmt.Printf("%#v\n", c1)
}

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDream(dreams []string) {
	// 修改切片将影响结构体内的值
	//p.dreams = dreams

	// 正确做法 先开辟空间 再复制
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

// StructPractice2 结构体 成员是引用类型  切片或map产生的问题
func StructPractice2() {
	p := &Person{
		name: "那一个人",
		age:  2,
	}
	d := []string{
		"gg",
		"hh",
		"en",
	}
	fmt.Println("p", p)
	p.SetDream(d)
	d[1] = "gg"
	fmt.Println(d)
	fmt.Println(p)
}

// CalcPractice 加减乘除
func CalcPractice() {
	c := &calc.Calc{Num: 2}
	fmt.Println(c.Sum(1))
	fmt.Println(c.Sub(2))
	fmt.Println(c.Multi(3))
	fmt.Println(c.Division(4))
}

func SlicePractice() {
	arr := [...]int{
		1, 2, 3, 54, 6, 7,
	}
	fmt.Printf("arr: %v v: %v t: %T", arr, arr[0:2:5], arr[:0:0])
	fmt.Println()
	// [low:high:max]
	// low 从...开始
	// high 到...
	// max 容量  max > high - low

	fmt.Println(cap(arr[0:2:5]))
}

// GoroutinePractice 20211018 练习goroutine
func GoroutinePractice() {
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			fmt.Println("i", i, 'i')
		}(i)
	}
	w.Wait()
}

func ChanPractice() {
	var w sync.WaitGroup
	w.Add(1)
	ch := make(chan int)
	go func() {
		ch <- 10

		w.Done()
	}()
	fmt.Println("ch", <-ch)
	w.Wait()
}

func SelectPractice() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		fmt.Println("当前迭代：", i)
		select {
		case x := <-ch:
			fmt.Println("取出", x)
		case ch <- i:
			fmt.Println("写入", i)
		}
	}
}

// StdinPractice 输入练习
func StdinPractice() {
	a := make([]byte, 6)
	n, err := os.Stdin.Read(a)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("n", n)
	fmt.Printf("%v %s %T", a, a, a)
}

// 多路复用  select配合channel
// *********************************
func launch() {
	fmt.Println("nuclear launch detected")
}

func commencingCountDown(canLunch chan int) {
	c := time.Tick(1 * time.Second)
	for countDown := 20; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<-c
	}
	canLunch <- -1
}

func isAbort(abort chan int) {
	os.Stdin.Read(make([]byte, 1))
	abort <- -1
}

func SelectChanPractice() {
	fmt.Println("Commencing countdown")

	abort := make(chan int)
	canLunch := make(chan int)
	go isAbort(abort)
	go commencingCountDown(canLunch)
	select {
	case <-canLunch:

	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

//**************************

func NetHttpPractice() {
	http.HandleFunc("/apiTest", func(writer http.ResponseWriter, request *http.Request) {
		data := "<h3>aa</h3>"
		_, err := fmt.Fprintln(writer, data)
		if err != nil {
			return
		}
		writer.Header().Set("content-type", "text/html")
	})
	err := http.ListenAndServe(":9191", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DownPic() {
	// 图片以花瓣网的图片为例
	//imgUrl := "http://hbimg.b0.upaiyun.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320"
	imgUrl := "https://qr.m.jd.com/show?appid=133&size=300&t="

	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	// defer后的为延时操作，通常用来释放相关变量
	defer res.Body.Close()

	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create("./test.png")
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	// 输出文件字节大小
	fmt.Printf("Total length: %d", written)
}

func JdSign() {
	//const url = "https://qr.m.jd.com/show?appid=133&size=300&t="
	const url = "https://www.liwenzhou.com/images/qrcode_for_gzh.jpg"
	//const redirectUrl = "https://passport.jd.com/new/login.aspx"
	req, _ := http.NewRequest("get", url+strconv.Itoa(int(1000*time.Now().Unix())), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Referer", redirectUrl)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	client := &http.Client{}
	resp, _ := client.Do(req)
	//reader := bufio.NewReaderSize(resp.Body, 32*1024)
	f, err := os.Create("./qr_code.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(f, resp.Body)
}
