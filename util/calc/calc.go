package calc

import "fmt"

// Calc 加减乘除
type Calc struct {
	Num int
}

func (c *Calc) Sum(num int) int {
	c.Num += num
	return c.Num
}

func (c *Calc) Sub(num int) int {
	return c.Num - num
}

func (c *Calc) Multi(num int) int {
	return c.Num * num
}

func (c *Calc) Division(num int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("\nSomething wrong!!")
			fmt.Println("Error", err)
		}

	}()
	if c.Num == 0 {
		panic("Division by zero is not allowed")
	}
	return c.Num / num
}
