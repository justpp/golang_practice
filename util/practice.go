package util

import "fmt"

func NNTable() {
	for j := 9;j>0;j--{
		for i:=1;i<j+1;i++ {
			fmt.Printf("%v*%v=%v;",j,i,j*i)
		}
		fmt.Println()
	}
}

func SwitchDemo()  {
	a:= 'a'
	n:=0
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
	fmt.Println("n:",n)
}
