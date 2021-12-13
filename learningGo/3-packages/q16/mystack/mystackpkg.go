package mystack

import "fmt"

type MyStack struct {
	Top   int
	Items [10]int
}

func New() *MyStack {
	return &MyStack{
		Top:   0,
		Items: [10]int{},
	}
}

func (ms *MyStack) Push(item int) {
	if ms.Top == cap(ms.Items) {
		fmt.Println("stack full")
		return
	}
	ms.Top++
	ms.Items[ms.Top] = item
}

func (ms *MyStack) Pop() int {
	poped := ms.Items[ms.Top]
	ms.Items[ms.Top] = 0
	ms.Top--

	return poped
}
