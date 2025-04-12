package main

import (
	"fmt"
	"io/ioutil"
	// "os"
)

func main() {
	// file,err:=os.Create("exmaple.txt")
	// if err!=nil{
	// 	fmt.Println("Error hile creating file",err)
	// 	return

	// }
	// defer file.Close()

	// content:="hello world by amith"
	

	// byte, error := io.WriteString(file,content+"\n")
	// fmt.Println("byte",byte)
	// if error !=nil {
	// 	fmt.Println("Error while writing the file",error)
	// 	return
		
	// }

	// fmt.Println("sucessfully craeted file")


	//read the data

	// file,err:=os.Open("exmaple.txt")
	// if err!=nil{
	// 	fmt.Println("errro while opening file")
	// 	return;
	// }
	// defer file.Close()

	// //crate the buffer to read the file

	// buffer:=make([]byte,1024)

	// for {
	// 	n,err:=file.Read(buffer)
	// 	if err ==io.EOF {
	// 		break
	// 	}
	// 	if err!=nil {
	// 		fmt.Println("Error whie reading the file",err)
	// 		return
	// 	}
	// 	//process the file
	// 	fmt.Println(string(buffer[:n]))
	// }


	content,err:=ioutil.ReadFile("exmaple.txt")
	if err!=nil {
		fmt.Println("error while raeding the file",err)
		return
		
	}
	fmt.Println(string(content))

	

}