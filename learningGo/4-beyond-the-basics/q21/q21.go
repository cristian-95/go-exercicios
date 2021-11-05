/*Q21. (1) Linked List
1. Make use of the package container/list to create a (doubly) linked list. Push the
values 1, 2 and 4 to the list and then print it.
2. Create your own linked list implementation. And perform the same actions as in
question 1	*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 1.
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("----------------------------")
	// 2.
	ml := New()
	ml.PushBack("a")
	ml.PushBack("b")
	ml.PushBack("c")
	for i := ml.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Content)
	}

}

// 2.
type List struct {
	Content      string
	PreviousNode *List
	NextNode     *List
}

func New() *List {
	return &List{
		Content:      "",
		PreviousNode: nil,
		NextNode:     nil,
	}
}

func (l *List) PushBack(content string) {
	if l.PreviousNode == nil && l.NextNode == nil {
		l.Content = content
		l.NextNode = New()
		l.PreviousNode = nil
	} else if l.NextNode != nil {
		l.NextNode.PushBack(content)
	} else {
		l.NextNode = &List{
			Content:      content,
			PreviousNode: l,
			NextNode:     nil,
		}
	}
}

func (l *List) Front() *List {
	if l.PreviousNode != nil {
		return l.PreviousNode.Front()
	} else {
		return l
	}
}

func (l *List) Next() *List {
	return l.NextNode
}
