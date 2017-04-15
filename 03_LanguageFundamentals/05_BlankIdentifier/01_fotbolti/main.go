package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http" // import the package with "get" function
)

func main() {
	res, err := http.Get("http://fotbolti.net/") //Error checking // Get returns a non-nil error number if it fails to open the page
	if err != nil {                              //Error Checking
		log.Fatal(err) //Error Checking
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil { // If "err" is not equal to "nil"
		log.Fatal(err) // display error message for "err"
	}
	fmt.Printf("%s", page) // Print variable "page" as a string
}
