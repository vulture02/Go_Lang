package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning web services")
	res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err !=nil{
		fmt.Println("Error getting reponse ",err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("type of response %T\n",res);
	
	
	
	
	//fmt.Println("response",res)//this response we cant ACCESSS THE DATA


	data,err:=ioutil.ReadAll(res.Body)
	if err!=nil {
		fmt.Println("Error reading response",err);
		return
		
	}
	fmt.Println("response",string(data))

}