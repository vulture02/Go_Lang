package main

import "fmt"
func modifyvaluerefrence( num *int){
	*num=*num+2
}

func main() {
	// var num int
	// num = 2
	// var ptr *int
	// ptr = &num
	// num:="amith"
	// ptr:=&num
	// // fmt.Println("numhas the value",num)
	// fmt.Println("pointer contain",ptr)
	// fmt.Println("data contain through pointer",*ptr)
	value:=10
	modifyvaluerefrence(&value)
	fmt.Println("value contain ",value)
}
