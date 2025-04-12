package main

import "fmt"

func main() {
	var fname, lname string
	
	fmt.Print("Enter your first name: ")
	fmt.Scan(&fname)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lname)

	fmt.Println("Your name is:", fname, lname)
}

