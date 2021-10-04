package main

import "fmt"

func average(slice []float64) (result float64) {
	if len(slice) == 0 {
		return
	}
	var sum float64
	for _, value := range slice {
		sum += value
	}
	result = sum / float64(len(slice))
	return
}

func main() {
	slc := []float64{1.00, 5.56, 9.09, 199.56, 23.2}
	num := average(slc)
	fmt.Println(num)

	num2 := average([]float64{})
	fmt.Println(num2)

}
