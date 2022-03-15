package practice

import (
	"fmt"
	"testing"
)

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")

	fmt.Printf("t1c%T ,%p \n", t1, t1)
	t2 := t1.Clone()
	fmt.Printf("t2c%T ,%p \n", t2, t2)
	if t1 == t2 {
		t.Fatal("error! get clone not working!")
	}
}

func TestCloneFormManager(t *testing.T) {
	c := manager.Get("t1").Clone()
	fmt.Printf("c%T ,%p \n", c, c)
	t1 := c.(*Type1)
	fmt.Printf("tcc%T ,%p \n", t1, t1)
	if t1.name != "Type1" {
		t.Fatal("error!")
	}
}

func init() {
	manager = NewPrototypeManager()
	t1 := &Type1{
		name: "Type1",
	}
	fmt.Printf("%T ,%p \n", t1, t1)
	manager.Set("t1", t1)
}
