package main

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent tasks in Go programs.
// Goroutines are created using the `go` keyword followed by a function call.
// The function will run concurrently with the rest of the program.
func syahello(){
	fmt.Println("Hello, Goroutines! frist")// This function is called syahello
	// time.Sleep(2000*time.Millisecond)// Sleep for 100 milliseconds
	fmt.Println("sleeping for 100 milliseconds second")// Print a message after sleeping
}
func hei(){
	fmt.Println("Hello,ANUSH! third")//
	time.Sleep(1000*time.Millisecond)// Sleep for 100 milliseconds
	fmt.Println("sleeping for 100000 milliseconds fourth")// Print a message after sleeping
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Learning, Goroutines!")
	go syahello()// This starts the goroutine
	// The main function continues to run concurrently with the goroutine.
	go hei()

	// Wait for the goroutine to finish before exiting the program.
	time.Sleep(900*time.Millisecond)// Sleep for 100 milliseconds
}
//we can use two goroutines at the same time
//but we have to use sleep function to wait goroutines to finish
//or else the main fuction will run for  two goroutines and exit the program
//if time is give less than in main function if upto 1000 millsecond if  in go rountine there is 1000 and anthor gorountine has 2000 it will
//excute only 1000 and exit the program	