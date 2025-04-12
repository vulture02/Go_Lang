package main

import (
	"fmt"
	"strings"
)

func main() {
	// data := "apple.orange.banana"
	// parrts := strings.Split(data, ".")//here spraeted based on , now it is in .
	// fmt.Println(parrts)
	// str:="one two three four two two five"
	// count:=strings.Count(str,"two");
	// fmt.Println(count)
	// str:="   hello ,        go!  "
	// trimed:=strings.TrimSpace(str)//it will trime olny at atrat and end of the strings
	// fmt.Println(trimed)

	str1:="amith"
	str2:="anish";
	resilt:=strings.Join([]string{str1,str2}," ")//this will make tem in array and add what is persnt in double qoute
	fmt.Println(resilt)


}