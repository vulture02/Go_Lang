package main

import "fmt"

func main() {
	//intinalize map 
	studentGrades := make(map[string]int)
	studentGrades["prince"] = 34
	studentGrades["amith"] = 99
	studentGrades["sushanth"] = 94
	studentGrades["sanjay"] = 85
	// fmt.Println("marks of amith is",studentGrades["amith"])
	// fmt.Println("marks of sudhanth",studentGrades["sushanth"])

	//cheching exists or not with boolean
	// grades,exists:=studentGrades["sanjay"]
	// fmt.Println("Grands of sanjay",grades);
	// fmt.Println("sanjay exists",exists)
	// fmt.Println("marks of snajay",studentGrades["sanjay"])

	// for index, value := range studentGrades {
	// 	fmt.Printf("key is %s and marks is %d\n", index, value)
	// }
	person:=map[string]int{
		"alice":90,
		"bob":85,
		"charlie":95,
	}
	for i,v:=range person{
		fmt.Printf("key and %s and marks is %d\n",i,v)
	}
	
}  
