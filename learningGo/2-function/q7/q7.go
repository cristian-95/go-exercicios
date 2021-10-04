package main

import "fmt"

func order(x, y int) (int, int) {
	if x <= y {
		return x, y
	} else {
		return y, x
	}

}

func main() {
	fmt.Println(order(82, 8))

}
