/*Q12. (1) Map function
A map()-function is a function that takes a function and a list.
The function is applied to each member in the list and a new list containing these
calculated values is returned.
Thus:
	 map(f (), (a1, a2, . . . , an−1, an)) = (f (a1), f (a2), . . . , f (an−1), f (an))
1. Write a simple map()-function in Go. It is sufficient for this function only to work
for ints.
2. Expand your code to also work on a list of strings.	*/
package main

import (
	"fmt"
	"strings"
)

// 1.
func simpleMapFunction(f func(int) int, data []int) (res []int) {
	for _, v := range data {
		res = append(res, f(v))
	}
	return
}

// 2.
func anotherMapFunction(f func(string) string, data []string) (res []string) {
	for _, v := range data {
		res = append(res, f(v))
	}
	return
}

func main() {
	myFun := func(x int) int {
		return x * 10
	}
	data := []int{1, 2, 3, 4, 5, 6, 99}
	res := simpleMapFunction(myFun, data)
	fmt.Println(res)

	myAnotherFun := func(s string) string {
		return strings.Join(
			[]string{
				strings.Repeat(s[:len(s)/2], 2),
				strings.Repeat(s[len(s)/2:], 2)},
			"")
	}
	anotherData := []string{"caju", "manga", "uva", "jaca"}
	anotherRes := anotherMapFunction(myAnotherFun, anotherData)
	fmt.Println(anotherRes)
}
