package util

import (
	"fmt"
	"os"
)

// Student 结构体 学生
type Student struct {
	Name string
	Age  int
}

type input interface {
	pickUp(i *input)
}

// StudentList 存储map
var StudentList = make(map[uint64]*Student, 12)

// Manger 学生管理系统
// func (s Student) Manger() {
// 	fmt.Println("23")
// }

// Manger 学生管理系统
func Manger() {
	for {
		StudentList[1] = &Student{"二傻子", 2}
		fmt.Print(`
		学生管理系统：
			1.查看所有学生
			2.新增学生
			3.删除学生
			4.退出
		请输入命令编号:`)
		var input int
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("输入的是个什么玩意儿···")
		}
		println("你输入了:", input)

		switch input {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
			showAllStudent()
		case 3:
			delStudent()
			showAllStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入的是个什么玩意儿···")
		}
	}

}

func createStudent(name string, age int) *Student {
	return &Student{name, age}
}

func showAllStudent() {
	fmt.Println("*******************************************")
	for k, v := range StudentList {
		fmt.Printf("学号:%v 姓名:%s 年龄:%v岁 \n", k, v.Name, v.Age)
	}
}

func addStudent() {
	var (
		id   uint64
		name string
		age  int
	)
	//p := &name
	name = "23"
	//gan(p)

	if id == 0 {
		fmt.Print("请输入学号:")
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Println("")
			fmt.Println("学号只支持数字！")
			addStudent()
		}
	}
	if name == "" {
		fmt.Print("请输入姓名:")
		_, er := fmt.Scanln(&name)
		if er != nil {
			fmt.Println("")
			fmt.Println("姓名啊，老弟！")
			addStudent()
		}
	}

	if age == 0 {
		fmt.Print("请输入年龄:")
		_, e := fmt.Scanln(&age)
		if e != nil {
			fmt.Println("")
			fmt.Println("年龄是数字啊，老弟！")
			addStudent()
		}
	}
	StudentList[id] = createStudent(name, age)
}

func delStudent() {
	var id uint64
	fmt.Print("请输入需要删除学号:")
	_, er := fmt.Scanln(&id)
	if er != nil {
		fmt.Println("")
		fmt.Println("学号啊，老弟！")
		addStudent()
	}
	delete(StudentList, id)
}

func gan(v input) {
	fmt.Println(&v)
}
