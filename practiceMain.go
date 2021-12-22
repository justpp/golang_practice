package main

import (
	"fmt"
	"giao/practice"
	"reflect"
)

type Type1 struct {
	Name string
}

func (t *Type1) Clone() practice.Cloneable {
	fmt.Printf("t %p \n", t)
	tc := *t
	fmt.Printf("tc %p \n", tc)
	fmt.Println("tc ln", tc)
	return &tc
}

func (t *Type1) SetName(name string) {
	t.Name = name
}

func main() {
	t1 := &Type1{
		Name: "T12",
	}

	t2 := t1.Clone()
	of := reflect.TypeOf(t2)
	fmt.Println("type of t2:", of)
	type1 := t2.(*Type1)
	type1.SetName("lll")
	fmt.Printf("t1f %p \n", t1)
	fmt.Printf("t2f %p \n", t2)
	fmt.Println("t2 ln", t2)

	fmt.Println("t2 ed ln", t2)
	fmt.Println("t1", t1)
}
