/*Q19. (2) Map function with interfaces
1. Use the answer from exercise Q12, but now make it generic using interfaces. Make
it at least work for ints and strings.*/
package main

import (
	"fmt"
)

type MyType struct {
	s string
	i int
}

//1. incompleto
func q19(f func(interface{}) string, data []MyType) (res []interface{}) {
	for _, v := range data {
		res = append(res, f(v))
	}
	return
}

func main() {
	fun1 := func(x interface{}) string {
		return fmt.Sprintf("%T", x)
	}
	data1 := []MyType{
		{"texto", 1},
		{"codigo", 9},
		{"carta", 99}}
	res1 := q19(fun1, data1)
	fmt.Println(res1...)

}
