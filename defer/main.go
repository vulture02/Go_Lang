package main

import "fmt"

func add(a,b int) int{
	return a+b;

}
func mul(a,b int) int{
	return a*b;
}

func main() {
	fmt.Println("starting of of program")
	ans:=mul(5,6);
	fmt.Println(ans);
	data:= add(5,6);
	defer fmt.Println(data)
	defer  fmt.Println("middel of  progarm")
	fmt.Println("end of progarm")
}
//defer key word excute the in last