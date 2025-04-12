package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning url")
	myurl:="https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"
	fmt.Printf("%T\n",myurl)

    parseurl,err:=url.Parse(myurl)
	if err!=nil {
		fmt.Println("cnat parse url ",err)
		return;
		
	}
	// fmt.Printf("Type of url %T",parseurl)
	fmt.Println("scheme",parseurl.Scheme)
	fmt.Println("hose",parseurl.Host)
	fmt.Println("path",parseurl.Path)
	fmt.Println("query",parseurl.RawQuery)

	//modifying the url componets
	parseurl.Path="/newPa"
	parseurl.RawQuery="username=go"

	//constructon a url string from a url object
	neurl:=parseurl.String()
	fmt.Println("new url",neurl)
}