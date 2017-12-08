package main

import (
	"fmt"
)

// I have an application, I have gold users.
// I want to print out names of the users,
// their gold points and also finish with (,GLD)

// User represents a gold user
type User struct {
	names  string
	points string
}

// Printer is a function that can be passed through other funcs
type Printer func(string)

// PrintDetails prints the name and points
func PrintDetails(u User, do Printer) {
	d := CreateDetails(u.names, u.points)
	do(d)
}

// CreateDetails returns a string of user details
func CreateDetails(names, points string) string {
	return names + " " + points
}

// Print does a fmt.Print()
func Print(s string) {
	fmt.Print(s)
}

// PrintLine prints each output in its own line
func PrintLine(s string) {
	fmt.Println(s)
}

//CreatePrinterFunc creates a customized Printer instance
func CreatePrinterFunc(s string) Printer {
	return func(str string) {
		fmt.Println(str + " " + s)
	}
}

func main() {
	user := User{"King Kunta", "901823"}
	PrintDetails(user, CreatePrinterFunc("GLD"))
}
