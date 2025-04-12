package main

import "fmt"

type person struct{
	FrisName string
	lasteNmae string
	Age int
}
type Contact struct{
	Email string
	phone string
}
type Address struct{
	House int
	Area string
	State string
}
type Employee struct{
	person_details person
	person_contact Contact
	person_adress Address

}
func main(){
	// var sanjay person
	// fmt.Println("sanjay peron :",sanjay)//out put is " "" "0 default value will print
	// sanjay.FrisName="N Sanjay"
	// sanjay.lasteNmae="Kumar"
	// sanjay.Age=20
	// fmt.Println(sanjay)

	// //another method(2)
	// person1 :=person{
	// 	FrisName: "akash",
	// 	lasteNmae: "mS",
	// 	Age: 20,
	// }
	// fmt.Println("person",person1)

	// //another method using new keyword
	// var person2=new(person)
	// person2.FrisName="amith"
	// person2.lasteNmae="polya"
	// person2.Age=20
	// fmt.Println("person 2:",person2.FrisName)//retriving particular data

	var  employee1 Employee
	employee1.person_details=person{
		FrisName: "amith",
		lasteNmae: "polya",
		Age: 20,
	}
	// another method
	employee1.person_contact.Email="amithp0210@gmail.com"
	employee1.person_contact.phone="7204213097"

	employee1.person_adress=Address{
		House: 55720,
		Area: "PUTTUR",
		State :"KARNATAKA",

	}
	// fmt.Println("EMPLOYEE1" ,employee1)
	fmt.Printf(employee1.person_adress.Area)

}

//structure insied structure