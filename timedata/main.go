package main

import (
	"fmt"
	"time"
)

func main(){
	//2006-01-02  15:04:05 this id fixed formate
	currentTime := time.Now()//built in function to get current time
	fmt.Println("\n",currentTime)
	fmt.Printf("type of currntTime %T\n",currentTime)


	formated:=currentTime.Format("02-01-2006,15:04:15")
	fmt.Println(formated)
	
	layout_str :="2006-01-02"
	datastr:="2023-11-25"
	formated_time,_:=time.Parse(layout_str,datastr)
	fmt.Println(formated_time)


	new_data:=currentTime.Add(24*time.Hour)//adding one day for current day
	fmt.Println("new_date time:",new_data)
	formated_new_data:=new_data.Format("2006/01/02 Monday")
	fmt.Println("formatted_new_data",formated_new_data)
	

}