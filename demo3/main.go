package main

import "fmt"

//数组栈
type ArrayStack struct {
	data  []int
	index int
	count int
}

//向数组栈顶添加一项
func (stack *ArrayStack) Push(item int) {
	stack.index++
	stack.count++
	stack.data[stack.index] = item
	return
}

//从数组栈顶删除
func (stack *ArrayStack) Pop() (v int) {
	v = stack.data[stack.index]
	stack.data[stack.index] = 0
	stack.index--
	stack.count--
	return
}

func main() {
	var arrayStack *ArrayStack
	arrayStack = &ArrayStack{
		data:  make([]int, 9),
		index: -1,
		count: 0,
	}

	arrayStack.Push(100)
	arrayStack.Push(88)
	arrayStack.Push(66)
	arrayStack.Push(55)
	arrayStack.Push(33)
	fmt.Println(arrayStack)

	arrayStack.Pop()

	fmt.Println(arrayStack)
}
