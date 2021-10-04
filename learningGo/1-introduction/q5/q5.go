/*Q5. (1) Average
1. Write code to calculate the average of a float64 slice. In a later exercise (Q6 you
will make it into a function */
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
