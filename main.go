package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int // int will default to bit of computer

	whatToSay = "Goodbye, cruel world."
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething() // set whatWasSaid to the type returned by saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) { // must return a string
	return "something", "else"
}
