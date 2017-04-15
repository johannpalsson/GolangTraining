package main

import "fmt"

const p string = "death & taxes" // string not necessary but works

const q = 42

func main() {
	//const q = 42 // also works inside a function
	const g string = "Hakuna Matata"
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	//fmt.Println(g , p, q) // runs and returns this: exit status 3221225506
	fmt.Println("Timon and Pumba say: ", g)
}

// a CONSTANT is a simple unchanging value
