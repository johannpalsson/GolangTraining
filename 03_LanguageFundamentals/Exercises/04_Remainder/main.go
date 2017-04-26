package main

import "fmt"

func main() {
	var smallnmbr int
	var largernmbr int
	fmt.Print("Enter a small number ")
	fmt.Scan(&smallnmbr)
	fmt.Print("Enter  a larger number ")
	fmt.Scan(&largernmbr)
	var remainder = largernmbr % smallnmbr
	fmt.Println("The remainder of ", largernmbr, " divided by ", smallnmbr, " is ", remainder)
}
