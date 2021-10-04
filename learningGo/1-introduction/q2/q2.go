package main

import "fmt"

func main() {
	//1.
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//2.
	// 	i := 0
	// LOOP:
	// 	fmt.Println(i)
	// 	i++
	// 	if i < 10 {
	// 		goto LOOP
	// 	}

	//3.
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)

}
