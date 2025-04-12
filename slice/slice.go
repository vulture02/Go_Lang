package main

import (
	"fmt"
	"slices"
)

func main(){
	// number:=[3]int{};
	// number=append(number,6,7,8,9,10)
	// fmt.Print(number)
	// fmt.Println("\n%T",number)
	// fmt.Println("leghth",len(number))
	// age:=[]string{}
	// fmt.Println(age)
	// fmt.Println("slice",number);
	// fmt.Println("length",len(number));
	// fmt.Println("capacity",cap(number));
	// numbers:=make([]int,3,5);
	// numbers=append(numbers, 6)
	// numbers=append(numbers, 96)
	// numbers=append(numbers, 9)
	// numbers=append(numbers, 6)
	// numbers=append(numbers, 96)
	// numbers=append(numbers, 9)
	// numbers=append(numbers, 9)
	// numbers=append(numbers, 96)
	//make a slice of length 3 and capacity 5
	// fmt.Println("slice",numbers);
	// fmt.Println("length",len(numbers));
	// fmt.Println("capacity",cap(numbers));
	// stock:=[]string{}
	// stock:=make([]int,0)
	// fmt.Println("slice",stock);
	// fmt.Println("length",len(stock));
	// fmt.Println("capacity",cap(stock));
	// prices:=[]int{10,20,30,40,50}
	// prices[2]=100
	// fmt.Println(prices)
	// prices=append(prices, 80,60,100,120)
	// fmt.Println(prices)
	numbers:=[]int{1,2,3,4,5}
	fmt.Println("slice",numbers);
	fmt.Println("length",len(numbers));
	fmt.Println("capacity",cap(numbers));
	dest := make([]int, len(numbers))
	copy(dest,numbers)
	fmt.Println("slice",dest);
	fmt.Println("length",len(dest));
	subslices:=numbers[1:4]
	fmt.Println(subslices)



}