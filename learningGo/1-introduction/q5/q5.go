package main

import "fmt"

func main() {
	slc := []float64{3.596, 1.00, 95.5, 20.55}
	var sum float64
	var avg float64
	for _, value := range slc {
		sum += value
	}
	avg = sum / float64(len(slc))
	fmt.Println(avg)

}
