package main

import "fmt"

func division(z int) (int, bool) { // takes int and assigns it to z
  // returns an int and a bool
  // could also use float64 to get decimals but then i need to convert z int to float 64
	return z / 2, z/2 == 0 // 1st return is z divided by 2 then we validate the bool
  // if using float64 then....  return float64(z) to convert
}
func main() {
	division := func (z int) (int, bool) { //assign the previous func to var in order to put it inside another func
		return z / 2, z/2 == 0 // 1st return is z divided by 2 then we validate the bool
	  // if using float64 then....  return float64(z) to convert
	}
	x, even := division(48) //assign my func division to two variables
  // put in a number, in this case 48
	fmt.Println(x, even) // print the value of the variables
}
