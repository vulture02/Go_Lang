package main

import(
	 "fmt" 
	 "sync" 
)
 func workers(i int, wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter //signal that the goroutine is done
	fmt.Printf("worker %d started\n",i)
	//do some work
	fmt.Printf("worker %d end\n",i)
	
 }
 //worker1
//worker2
//worker3

func main() {
	fmt.Println("explore gorountine statred")

	//start 3 workers goroutines
	var wg sync.WaitGroup
	for i := 1; i <=3; i++ {
		wg.Add(1) //increment the counter
	    go workers(i,&wg);
	}
	wg.Wait() //wait for all goroutines to finish
	//wait for all goroutines to finish before main exits
	fmt.Println("worker completed")
	//it does not wait for goroutines to finish main is bother aboout its function only
	// so we need to use sync.WaitGroup to wait for all goroutines to finish
}
