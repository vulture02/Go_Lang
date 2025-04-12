package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`        // lowercase "id" to match JSON
	Title     string `json:"title"`     // lowercase string
	Completed bool   `json:"completed"` // keep as is
}

func performRequest(){
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // fetching one item
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting response:", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }

	// fmt.Println("Data:", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo) // decode JSON into struct
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Todo:", todo)

}
func performPostRequest(){
	 todo:= Todo{
		UserId:123,
		Title:"Amith",
		Completed:true,

	}
	//convert the todo struct to json

	jsonData,err:=json.Marshal(todo)
	if err!=nil{
		fmt.Println("Error marshaling:",err)
		return
	}

	myurl:="https://jsonplaceholder.typicode.com/todos"
	//set the content type to application/json
	jsonString:=string(jsonData)
	fmt.Println("JSON String:",jsonString)
// create a new reader from the json string
    jsonreadre:=strings.NewReader(jsonString)

	//post the data to the url
	//http.Post(url,content-type,body)
	res,err:=http.Post(myurl,"application/json",jsonreadre)
	if err!=nil{
		fmt.Println("Error posting:",err)
		return
	}
	// check the status code of the response
	defer res.Body.Close()// close the response body when done
	fmt.Println("Response Status:",res.Status)
  // read the response body
	data,err:=ioutil.ReadAll(res.Body)
	if err!=nil{ 
		fmt.Println("Error reading response body:",err)
		return
	}
	fmt.Println("Response Body:",string(data))
    
}

func main() {
	fmt.Println("CURD")
	performRequest()
	performPostRequest()

	
}