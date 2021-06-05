package stack

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (as *ArrayStack) Push(value interface{}) {
	as.top += 1

	if as.top > len(as.data)-1 {
		as.data = append(as.data, value)
	} else {
		as.data[as.top] = value
	}

}

func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}

	v := as.data[as.top]
	as.top--

	return v
}

func (as *ArrayStack) IsEmpty() bool {
	return as.top < 0
}

func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}

	return as.data[as.top]
}

func (as *ArrayStack) Flush() {
	as.top = -1
}

func (as *ArrayStack) Print() {
	if as.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := as.top; i >= 0; i-- {
			fmt.Println(as.data[i])
			fmt.Println(" ")
		}
	}

}
