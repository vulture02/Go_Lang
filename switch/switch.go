package main
import "fmt"
func main(){
	// day:=3
	// switch day{
	// case 1:
	// 	fmt.Println("Monday")		
	// case 2:
	// 	fmt.Println("Tuesday")	
	// case 3:
	// 		fmt.Println("Wednesday")
	// case 4:	
	// 		fmt.Println("Thursday")
	// default:
	// 		fmt.Println("Invalid day")
	// }
	// month:="jan"
	// switch month{
	// 	case"jan","feb","mar":
	// 		fmt.Println("First quarter")
	// 	case"apr","may","jun":
	// 		fmt.Println("Second quarter")
	// 	case"jul","aug","sep":
	// 		fmt.Println("Third quarter")
	// 	case"oct","nov","dec":
	// 		fmt.Println("Fourth quarter")
	// 	default:
	// 		fmt.Println("Invalid month")
	// }
	temperature:=30	
	switch temperature{
	case temperature<0:
		fmt.Println("Freezing")
	case temperature>=0 && temperature<=10:
		fmt.Println("Cold")
	case temperature>10 && temperature<=20:
		fmt.Println("Cool")
	case temperature>20 && temperature<=30:
		fmt.Println("Warm")
	case temperature>30 && temperature<=40:
		fmt.Println("Hot")
	default:
		fmt.Println("Very Hot")
	}
}