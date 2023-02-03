package practice

import (
	"bufio"
	"container/list"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"giao/util/calc"
	_ "github.com/go-sql-driver/mysql"
	errors "github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"runtime/trace"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

// PracticeClosure 闭包函数作用域练习
func PracticeClosure() {
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

// PracticeStruct1 结构体练习
func PracticeStruct1() {
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

// Stud 学生
type Stud struct {
	ID     int `json:"id"`
	Gender string
	Name   string
}

// PracticeStructJson 结构体转json练习
func PracticeStructJson() {
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
	// p.dreams = dreams

	// 正确做法 先开辟空间 再复制
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

// PracticeStruct2 结构体 成员是引用类型  切片或map产生的问题
func PracticeStruct2() {
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

// PracticeCalc 加减乘除
func PracticeCalc() {
	c := &calc.Calc{Num: 2}
	fmt.Println(c.Sum(1))
	fmt.Println(c.Sub(2))
	fmt.Println(c.Multi(3))
	fmt.Println(c.Division(4))
}

func PracticeSlice() {
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

// PracticeGoroutine 20211018 练习goroutine
func PracticeGoroutine() {
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

func PracticeChan() {
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

func PracticeSelect() {
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

// PracticeStdin 输入练习
func PracticeStdin() {
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

func PracticeSelectChan() {
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

// **************************

func PracticeNetHttp() {
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
	// imgUrl := "http://hbimg.b0.upaiyun.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320"
	imgUrl := "https://www.liwenzhou.com/images/qrcode_for_gzh.jpg"
	// imgUrl := "https://qr.m.jd.com/show?appid=133&size=300&t="

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

func DownLoadImg(body io.Reader, fileName string) {
	f, err := os.Create(fileName) //  ./qr_code.png
	if err != nil {
		panic(err)
	}
	readAll, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	write, err := f.Write(readAll)
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("length", write)
}

func add(x *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		*x = *x + 1
		lock.Unlock()
	}
	wg.Done()
}

func PracticeGo() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	var x int
	var lock = new(sync.Mutex)

	go add(&x, wg, lock)
	go add(&x, wg, lock)
	wg.Wait()
	fmt.Println(x)
}

// PracticeSelect2 示例多路复用
func PracticeSelect2() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("out", x)
		case ch <- i:
			fmt.Println("in", i)
		}
	}
}

func sliceAdd(s []int) {
	s3 := s
	s3 = append(s3, 0)
	s = append(s, s3...)
	for i := range s {
		s[i]++
	}
	fmt.Println("s_add", s)
}

func PracticeSlice2() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	sliceAdd(s1)
	sliceAdd(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}

func PracticeChan2() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	c := make(chan int, 1)
	go func(a, b int, c chan int) {
		sum := a + b
		c <- sum
		fmt.Println(sum)
		wg.Done()
	}(7, 2, c)
	go func(a, b int, c chan int) {
		sum := a + b
		c <- sum
		fmt.Println(sum)
		wg.Done()
	}(-8, 5, c)
	x, y := <-c, <-c
	fmt.Println(x + y)
	wg.Wait()
}

func PracticeFlag() {
	var name string
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.Parse()
	fmt.Println(name)
}

func PracticeLog() {
	logFile, err := os.OpenFile("./log/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.Println("giaogiao")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[justpp]")
	log.Println("这是一条很普通的日志。")
}

func PracticeFile() {
	file, err := os.Open("./application.ini")
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	var content []byte
	var tmp = make([]byte, 128)
	for {
		read, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("读完了")
			break
		}
		if err != nil {
			return
		}
		content = append(content, tmp[:read]...)
	}
	fmt.Println(string(content))
}

func Content1() {
	wg := sync.WaitGroup{}
	a := 0
	wg.Add(1)
	go func() {
		for {
			fmt.Println("worker")
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	for {
		if a > 10 {
			break
		}
		fmt.Println("wait", a)
		time.Sleep(time.Second)
		a++
	}
	wg.Wait()
	fmt.Println("over")
}

func Content2var() {
	wg := sync.WaitGroup{}
	exit := false
	a := 0
	wg.Add(1)
	go func() {
		for {
			fmt.Println("worker")
			time.Sleep(time.Second)
			if exit {
				break
			}
		}
		wg.Done()
	}()

	fmt.Println("over")
	for {
		if a > 3 {
			exit = true
		}
		if a > 10 {
			break
		}
		fmt.Println("wait", a)
		time.Sleep(time.Second)
		a++
	}
	wg.Wait()
}

func Content3chan() {
	wg := sync.WaitGroup{}
	a := 0
	c := make(chan struct{})

	wg.Add(1)
	go func(c chan struct{}) {
	LOOP:
		for {
			fmt.Println("worker")
			time.Sleep(time.Second)
			select {
			case <-c:
				break LOOP
			}
		}
		wg.Done()
		fmt.Println("下班了")
	}(c)

	for {
		if a == 3 {
			c <- struct{}{}
			close(c)
		}
		if a > 10 {
			break
		}
		fmt.Println("wait", a)
		time.Sleep(time.Second)
		a++
	}
	wg.Wait()
	fmt.Println("over")
}

func Content4sync() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	ctx, cancelFunc := context.WithCancel(context.Background())
	go func(c context.Context) {
		go func() {
		LOOP:
			for {
				fmt.Println("worker2")
				time.Sleep(time.Second)
				select {
				case <-c.Done():
					break LOOP
				}
			}
		}()
	LOOP:
		for {
			fmt.Println("worker1")
			time.Sleep(time.Second)
			select {
			case <-c.Done():
				break LOOP
			}
		}
		wg.Done()
	}(ctx)
	time.Sleep(time.Second * 3)
	cancelFunc()
	wg.Wait()
	fmt.Println("over")
}

type User struct {
	Id   int
	Name *string
}

//	type Stringer interface {
//	   String() string
//	} 在打印结构体内部变量是指针类型时会自动调用结构体的String方法
//
// func (u User) String 可以被打印变量自动调用
// func (u *User) String 可以被打印指针变量调用
func (u *User) String() string {
	return fmt.Sprintf("ID:%v name:%v\n", u.Id, *u.Name)
}

func StringPtr() {
	name := "justpp"
	user := &User{
		Id:   2,
		Name: &name,
	}
	fmt.Println(user)
}

// SelectPractice select 练习
func SelectPractice() {
	// 创建管道
	output := make(chan string, 10)
	// 子协程写数据
	go func(ch chan string) {
		count := 0
		for {
			select {
			case ch <- "hello":
				fmt.Println("write hello", count)
			default:
				fmt.Println("channel full")
			}
			time.Sleep(time.Millisecond * 500)
			count++
		}
	}(output)
	// 取数据
	for s := range output {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func ForPractice() {
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}
	w := sync.WaitGroup{}
	for _, url := range urls {
		// 在for内开启协程，协程内使的外部变量：重新赋值或直接传入 不然会导致变量是同一个值
		url := url
		w.Add(1)
		go func() {
			println("url", url)
			w.Done()
		}()
	}
	w.Wait()
}

func fmtReflectType(i interface{}) {
	r := reflect.TypeOf(i)
	fmt.Printf("name: %v kind: %v\n", r.Name(), r.Kind())
}

type myType int64

func ReflectPractice() {
	var a *float32
	var b myType
	var c rune
	type student struct {
		Name string
		Age  int8 `json:"age"`
	}
	stu := student{"justpp", 2}
	fmtReflectType(a)
	fmtReflectType(b)
	fmtReflectType(c)
	b = 2
	bT := reflect.ValueOf(&b)
	bT.Elem().SetInt(9)
	fmt.Println("b", b)
	fmtReflectType(stu)

	stdRef := reflect.TypeOf(stu)
	age := stdRef.Field(1)
	ageVal := reflect.ValueOf(age)
	fmt.Printf("name: %v index: %v json tag:%v val: %v", age.Name, age.Index, age.Tag.Get("json"), ageVal)
}

func FlagPractice() {
	var name string
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	fmt.Printf("args: %v \n", args)
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "golang", "帮助信息 --go")
		_ = goCmd.Parse(args[1:])
	case "php":
		var a string
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "name", "php", "帮助信息 --php")
		phpCmd.StringVar(&a, "a", "a", "帮助信息 --php a")
		_ = phpCmd.Parse(args[1:])
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name %v", name)
}

func GoAddPractice() {
	var a int32
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a++
		}()
	}
	wg.Wait()
	spendTime := time.Now().Sub(start).Nanoseconds()

	fmt.Printf("use nothing add result:%d, spend:%v\n", a, spendTime)
}

func MutexAddPractice() {
	var a int32
	var wg sync.WaitGroup
	var u sync.Mutex
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			u.Lock()
			a += 1
			u.Unlock()
		}()
	}
	wg.Wait()
	spendTime := time.Now().Sub(start).Nanoseconds()

	fmt.Printf("use mutex add result:%d, spend:%v\n", a, spendTime)
}

func AtomicAddPractice() {
	var a int32
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()
	}
	wg.Wait()
	spendTime := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use atomic add result:%d, spend:%v\n", a, spendTime)
}

type Menu struct {
	Id   int    `json:"id"`
	Pid  int    `json:"pid"`
	Name string `json:"name"`
}

type ResponseMenu struct {
	Name     string         `json:"name"`
	Children []ResponseMenu `json:"children"`
}

func GetDepartmentData() []*Menu {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3308)/vpea_erp_local")
	if err != nil {
		fmt.Println("err open:", err)
		return nil
	}
	query, err := db.Query("select * from departments")
	if err != nil {
		fmt.Println("err query:", err)
		return nil
	}
	var menus []*Menu
	columns, _ := query.Columns()
	// 每个列的值,将值获取到[]byte中
	vals := make([][]byte, len(columns))
	// query.Scan的参数,
	scans := make([]interface{}, len(columns))
	// 让每一行的数据都填充到[][]byte
	for k := range columns {
		scans[k] = &vals[k]
	}
	for query.Next() {
		query.Scan(scans...)
		row := make(map[string]string)
		for k, v := range vals {
			key := columns[k]
			row[key] = string(v)
		}
		menu := &Menu{}
		if err != nil {
			return nil
		}
		// 通过反射向结构体赋值
		menuRefT := reflect.TypeOf(menu).Elem()
		menuRefV := reflect.ValueOf(menu).Elem()
		for i := 0; i < menuRefV.NumField(); i++ {
			// 根据结构体tag找到对应字段的值
			rowVal := row[menuRefT.Field(i).Tag.Get("json")]
			vType := menuRefV.Field(i).Kind()
			switch vType {
			case reflect.Int:
				intVal, _ := strconv.Atoi(rowVal)
				menuRefV.Field(i).Set(reflect.ValueOf(intVal))
				break
			case reflect.String:
				menuRefV.Field(i).Set(reflect.ValueOf(rowVal))
			default:
				fmt.Printf("暂不支持类型:%s\n", vType.String())
			}
		}
		menus = append(menus, menu)
		if err != nil {
			fmt.Println("err scan:", err)
			return nil
		}
	}
	return menus
}

func CreateLoop(menus []*Menu, pid int) []ResponseMenu {
	tree := make([]ResponseMenu, 0)
	for _, menu := range menus {
		if menu.Pid == pid {
			tree = append(tree, ResponseMenu{
				Name:     menu.Name,
				Children: CreateLoop(menus, menu.Id),
			})
		}
	}
	return tree
}

// HeapSortMax arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 // 二叉深度
	for i := depth; i >= 0; i-- {
		topMax := i
		leftChild := i*2 + 1
		rightChild := i*2 + 2
		if leftChild <= length-1 && arr[leftChild] > arr[topMax] { // 防止越过界
			topMax = leftChild
		}
		if rightChild <= length-1 && arr[rightChild] > arr[topMax] { // 防止越界
			topMax = rightChild
		}
		if topMax != i {
			arr[i], arr[topMax] = arr[topMax], arr[i]
		}
	}
	return arr
}

// HeapSort 堆排序 最大堆 升序
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastIn := length - i
		HeapSortMax(arr, lastIn)
		if i < length {
			arr[0], arr[lastIn-1] = arr[lastIn-1], arr[0]
		}
	}
	return arr
}

func Unmarshal(data []byte, v interface{}) error {
	s := string(data)
	// 去除前后的连续空格
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	if len(s) == 0 {
		return nil
	}
	typ := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	if typ.Kind() != reflect.Ptr { // 因为要修改v，必须传指针
		return errors.New("must pass pointer parameter")
	}

	typ = typ.Elem() // 解析指针
	value = value.Elem()

	switch typ.Kind() {
	case reflect.String:
		if s[0] == '"' && s[len(s)-1] == '"' {
			value.SetString(s[1 : len(s)-1]) // 去除前后的""
		} else {
			return fmt.Errorf("invalid json part: %s", s)
		}
	case reflect.Bool:
		if b, err := strconv.ParseBool(s); err == nil {
			value.SetBool(b)
		} else {
			return err
		}
	case reflect.Float32,
		reflect.Float64:
		if f, err := strconv.ParseFloat(s, 64); err != nil {
			return err
		} else {
			value.SetFloat(f) // 通过reflect.Value修改原始数据的值
		}
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		if i, err := strconv.ParseInt(s, 10, 64); err != nil {
			return err
		} else {
			value.SetInt(i) // 有符号整型通过SetInt
		}
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		if i, err := strconv.ParseUint(s, 10, 64); err != nil {
			return err
		} else {
			value.SetUint(i) // 无符号整型需要通过SetUint
		}
	case reflect.Slice:
		if s[0] == '[' && s[len(s)-1] == ']' {
			arr := SplitJson(s[1 : len(s)-1]) // 去除前后的[]
			if len(arr) > 0 {
				slice := reflect.ValueOf(v).Elem()                    // 别忘了，v是指针
				slice.Set(reflect.MakeSlice(typ, len(arr), len(arr))) // 通过反射创建slice
				for i := 0; i < len(arr); i++ {
					eleValue := slice.Index(i)
					eleType := eleValue.Type()
					if eleType.Kind() != reflect.Ptr {
						eleValue = eleValue.Addr()
					}
					if err := Unmarshal([]byte(arr[i]), eleValue.Interface()); err != nil {
						return err
					}
				}
			}
		} else if s != "null" {
			return fmt.Errorf("invalid json part: %s", s)
		}
	case reflect.Map:
		if s[0] == '{' && s[len(s)-1] == '}' {
			arr := SplitJson(s[1 : len(s)-1]) // 去除前后的{}
			if len(arr) > 0 {
				mapValue := reflect.ValueOf(v).Elem()                // 别忘了，v是指针
				mapValue.Set(reflect.MakeMapWithSize(typ, len(arr))) // 通过反射创建map
				kType := typ.Key()                                   // 获取map的key的Type
				vType := typ.Elem()                                  // 获取map的value的Type
				for i := 0; i < len(arr); i++ {
					brr := strings.Split(arr[i], ":")
					if len(brr) != 2 {
						return fmt.Errorf("invalid json part: %s", arr[i])
					}

					kValue := reflect.New(kType) // 根据Type创建指针型的Value
					if err := Unmarshal([]byte(brr[0]), kValue.Interface()); err != nil {
						return err
					}
					vValue := reflect.New(vType) // 根据Type创建指针型的Value
					if err := Unmarshal([]byte(brr[1]), vValue.Interface()); err != nil {
						return err
					}
					mapValue.SetMapIndex(kValue.Elem(), vValue.Elem()) // 往map里面赋值
				}
			}
		} else if s != "null" {
			return fmt.Errorf("invalid json part: %s", s)
		}
	case reflect.Struct:
		if s[0] == '{' && s[len(s)-1] == '}' {
			arr := SplitJson(s[1 : len(s)-1])
			if len(arr) > 0 {
				fieldCount := typ.NumField()
				// 建立json tag到FieldName的映射关系
				tag2Field := make(map[string]string, fieldCount)
				for i := 0; i < fieldCount; i++ {
					fieldType := typ.Field(i)
					name := fieldType.Name
					if len(fieldType.Tag.Get("json")) > 0 {
						name = fieldType.Tag.Get("json")
					}
					tag2Field[name] = fieldType.Name
				}

				for _, ele := range arr {
					brr := strings.SplitN(ele, ":", 2) // json的value里可能存在嵌套，所以用:分隔时限定个数为2
					if len(brr) == 2 {
						tag := strings.Trim(brr[0], " ")
						if tag[0] == '"' && tag[len(tag)-1] == '"' { // json的key肯定是带""的
							tag = tag[1 : len(tag)-1]                        // 去除json key前后的""
							if fieldName, exists := tag2Field[tag]; exists { // 根据json key(即json tag)找到对应的FieldName
								fieldValue := value.FieldByName(fieldName)
								fieldType := fieldValue.Type()
								if fieldType.Kind() != reflect.Ptr {
									// 如果内嵌不是指针，则声明时已经用0值初始化了，此处只需要根据json改写它的值
									fieldValue = fieldValue.Addr()                                            // 确保fieldValue指向指针类型，因为接下来要把fieldValue传给Unmarshal
									if err := Unmarshal([]byte(brr[1]), fieldValue.Interface()); err != nil { // 递归调用Unmarshal，给fieldValue的底层数据赋值
										return err
									}
								} else {
									// 如果内嵌的是指针，则需要通过New()创建一个实例(申请内存空间)。不能给New()传指针型的Type，所以调一下Elem()
									newValue := reflect.New(fieldType.Elem())                               // newValue代表的是指针
									if err := Unmarshal([]byte(brr[1]), newValue.Interface()); err != nil { // 递归调用Unmarshal，给fieldValue的底层数据赋值
										return err
									}
									value.FieldByName(fieldName).Set(newValue) // 把newValue赋给value的Field
								}

							} else {
								fmt.Printf("字段%s找不到\n", tag)
							}
						} else {
							return fmt.Errorf("invalid json part: %s", tag)
						}
					} else {
						return fmt.Errorf("invalid json part: %s", ele)
					}
				}
			}
		} else if s != "null" {
			return fmt.Errorf("invalid json part: %s", s)
		}
	default:
		fmt.Printf("暂不支持类型:%s\n", typ.Kind().String())
	}
	return nil
}

// SplitJson 由于json字符串里存在{}[]等嵌套情况，直接按,分隔是不合适的
func SplitJson(json string) []string {
	rect := make([]string, 0, 10)
	stack := list.New() // list是双端队列，用它来模拟栈
	beginIndex := 0
	for i, r := range json {
		if r == rune('{') || r == rune('[') {
			stack.PushBack(struct{}{}) // 我们不关心栈里是什么，只关心栈里有没有元素
		} else if r == rune('}') || r == rune(']') {
			ele := stack.Back()
			if ele != nil {
				stack.Remove(ele) // 删除栈顶元素
			}
		} else if r == rune(',') {
			if stack.Len() == 0 { // 栈为空时才可以按,分隔
				rect = append(rect, json[beginIndex:i])
				beginIndex = i + 1
			}
		}
	}
	rect = append(rect, json[beginIndex:])
	return rect
}

// BinSearch 二分查找 数组得是升序
func BinSearch(arr []int, findData int) int {
	i := 0
	low := 0
	high := len(arr) - 1
	for low <= high {
		i++
		mid := int(uint(low+high) >> 1)
		if arr[mid] > findData {
			high = mid - 1
		} else if arr[mid] < findData {
			low = mid + 1
		} else {
			fmt.Printf("查找了%v次\n", i)
			return mid
		}
	}
	return -1
}

// 实现set
type Inner interface{}
type Set struct {
	m map[Inner]bool
	*sync.RWMutex
}

func (s *Set) New() *Set {
	return &Set{
		m: map[Inner]bool{},
	}
}

func (s *Set) Add(param Inner) {
	s.Lock()
	defer s.Unlock()
	s.m[param] = true
}

func Trace() {
	f, err := os.Create("./log/trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	fmt.Println("gg")
}

// Struct1 同时赋值的struct，直接比较地址是不相同的
// 若使用fmt.Printf("%p\n%p\n"后，猜测编译优化，值相等
func Struct1() {
	var a, b struct{}
	print(&a, "\n", &b, "\n")
	fmt.Printf("%p\n%p\n", &a, &b)
	fmt.Println(&a == &b)
}

func Struct2() {
	var c, d struct{}
	fmt.Printf("%p\n%p\n", &c, &d)
	print(&c, "\n", &c, "\n")
	fmt.Println(&c == &d)
}

func Goroutine1() {
	// c := make(chan os.Signal)
	var wg sync.WaitGroup
	start := time.Now()
	defer func() {
		fmt.Println("耗时:", time.Now().Sub(start).Seconds())
	}()
	for i := 1; i < 10000; i++ {
		wg.Add(1)
		go Goroutine1Attr1(&wg, i)
	}
	// signal.Notify(c, os.Interrupt)
	// select {
	// case <-c:
	// 	fmt.Println("out")
	// }
	wg.Wait()

}

func Goroutine1Attr1(wg *sync.WaitGroup, id int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()
	get, err := http.Get("http://127.0.0.1:999/" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer get.Body.Close()
	all, err := io.ReadAll(get.Body)
	if err != nil {
		panic(err)
	}
	u := struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}{}
	err = json.Unmarshal(all, &u)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}

// Mul1To00 1到100乘
func Mul1To00() {
	var sum = big.NewInt(1)
	// sum = sum.MulRange(1, 100)
	// 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000
	// 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000
	for i := 1; i <= 100; i++ {
		sum.Mul(sum, big.NewInt(int64(i)))
		fmt.Println(i, sum)
	}
}
