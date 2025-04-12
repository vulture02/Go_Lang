package main

import "fmt"

func main() {
	// for i:= 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	//infinete lopp
	// counter:=0
	// for{
	// 	fmt.Println("infinter loo")
	// 	counter++
	// 	if counter==3 {
	// 		break;
			
	// 	}
	// }
	number:=[]int{1,2,3,4,5};
	for i ,v:=range number{
		fmt.Printf("idex is %d ,and value is%d\n",i,v);
	}
	data:="hello world";
	for i,char:=range data{
		fmt.Printf("index of data is%d ,and value is%c\n",i,char)
	}

}