/*Q14. (1) Bubble sort
1. Write a function that performs a bubble sort on a slice of ints. From [24]:
It works by repeatedly stepping through the list to be sorted, comparing
each pair of adjacent items and swapping them if they are in the wrong
order. The pass through the list is repeated until no swaps are needed,
which indicates that the list is sorted. The algorithm gets its name from
the way smaller elements “bubble” to the top of the list.

[24] also gives an example in pseudo code:
procedure bubbleSort( A : list of sortable items )
	do
		swapped = false
		for each i in 1 to length(A) - 1 inclusive do:
			if A[i-1] > A[i] then
				swap( A[i-1], A[i] )
				swapped = true
			end if
		end for
	while swapped
end procedure
*/
package main

import "fmt"

func main() {
	numbers := []int{5, 9, 1, 10, 2, 4, 7, 3, 8, 6, 0}
	fmt.Printf("antes:\t%v\n", numbers)
	bubbleSort(numbers)
	//alternativa: sort.Ints(numbers)
	fmt.Printf("depois:\t%v\n", numbers)
}

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
}
