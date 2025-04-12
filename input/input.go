package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hey what is your name")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("hello mr",name)

	reader:=bufio.NewReader(os.Stdin) //reader space also entier name
	name,_:=reader.ReadString('\n')
	fmt.Println("hello mr",name)
	
}