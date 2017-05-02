package main

import "fmt"

func max(numbers ...int) int {
	var largest int             // initialize largest to its zero state.. so largest == 0
	for _, v := range numbers { // v is the numbers in the slice, individually
		if v > largest {
			largest = v // if v is larger than the individual number, it reassigns largest
		}
	}
	return largest
}
func main() {
	HighestNumber := max(1, 2, 45, 555, 66, 3333, 643, 23423, 3344) // assign func max as a var
	fmt.Println(HighestNumber)                                       //print the var
}
