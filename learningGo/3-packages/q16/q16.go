package main

import (
	"fmt"

	"github.com/go-exercicios/go-exercicios/learningGo/3-packages/q16/mystack"
)

func main() {
	mystk := mystack.New()

	mystk.Push(2)
	fmt.Println(mystk)

}
