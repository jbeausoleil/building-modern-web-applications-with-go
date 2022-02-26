package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString) // pointer points to a specific address in memory
	log.Println("After func call myString is set to", myString)

	var a = 45
	var s *int = &a
	log.Println(s)
}

func changeUsingPointer(s *string) { // type pointer to a string
	newValue := "Red"
	*s = newValue // find the memory address and modify its value
}
