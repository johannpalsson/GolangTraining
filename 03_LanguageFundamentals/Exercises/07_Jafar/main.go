package main

import "fmt"

func main() {
	counter := 0 //skilgreini var sem geymir gildi i
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i //ef tala er deilanleg með 3 þá bæti ég henni við i
		} else if i%5 == 0 {
			counter += i // same as above nema ef deilt með 5
		}
	}
	fmt.Println(counter)
}
