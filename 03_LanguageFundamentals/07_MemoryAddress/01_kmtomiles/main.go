package main

import "fmt"

const milestokm float64 = 1.609344

func main() {
	var mile float64
	fmt.Print("Enter distance in miles: ")
	fmt.Scan(&mile)
	km := mile * milestokm
	fmt.Println(mile, " miles is equal to: ", km, " kilometres!")
}
// actually calculates miles to kilometres, not vice versa ;)
