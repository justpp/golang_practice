package practice

import "fmt"

// 工厂接口
type operatorFactory interface {
	create() mathFactory
}

// 实例接口
type mathFactory interface {
	setOperatorA(int)
	setOperatorB(int)
	computedResult() int
}

// 实例基类
type baseOperator struct {
	operatorA, operatorB int
}

func (o *baseOperator) setOperatorA(operand int) {
	o.operatorA = operand
}

func (o *baseOperator) setOperatorB(operand int) {
	o.operatorB = operand
}

type plusOperator struct {
	*baseOperator
}

func (o *plusOperator) computedResult() int {
	return o.operatorA + o.operatorB
}

type multiOperator struct {
	*baseOperator
}

func (o *multiOperator) computedResult() int {
	return o.operatorA * o.operatorB
}

type plusFactory struct{}

func (o *plusFactory) create() mathFactory {
	// 嵌套结构体使用&实例化时，内部嵌套的也需要实例化，不然会空指针 nil pointer
	return &plusOperator{&baseOperator{}}
}

type multiFactory struct{}

func (o *multiFactory) create() mathFactory {
	return &multiOperator{&baseOperator{}}
}

// TestFactoryFunc 工厂方法
func TestFactoryFunc() {
	var factory operatorFactory
	var mathFactory mathFactory

	factory = &plusFactory{}
	mathFactory = factory.create()
	mathFactory.setOperatorA(2)
	mathFactory.setOperatorB(3)
	fmt.Println("plus:", mathFactory.computedResult())

	factory = &multiFactory{}
	mathFactory = factory.create()
	mathFactory.setOperatorA(2)
	mathFactory.setOperatorB(3)
	fmt.Println("multi:", mathFactory.computedResult())
}

// 抽象工厂 就是在工厂方法中可以产出同系列的实列
// 简单来说就是在工厂方法上嵌套一层工厂方法
// 比如
// 电子厂 产出 手机、电视
// 接小米的单 产出 小米手机、小米电视
// 接蓝绿的单 产出 蓝绿手机、蓝绿电视
