package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var num int = 42
	// // fmt.Println(num)
	// // fmt.Printf("type of num is %T",num)

	// var data float64=float64(num)
	// data=data+1.235
	// fmt.Println(data)
	// fmt.Printf("data type of data is%T",data)
	// num:=123
	// str:=strconv.Itoa(num);
	// fmt.Println("str is ",str)
	// fmt.Printf("%T",str)

	// number_str:="123"
	// number_int,_:=strconv.Atoi(number_str)//convert the string to the integer we should mention Atoi
	// fmt.Println("str is",number_str)
	// fmt.Printf("type of number is %T\n",number_int)

	number_string:="2.34";
	number_float,_:=strconv.ParseFloat(number_string,64)//while conversion of string to the folat we should use parse and which byte we should specifie
	fmt.Println("number is float",number_float)
	fmt.Printf("type of number is %T",number_float)
}