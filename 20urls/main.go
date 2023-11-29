package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:5000/learn?coursename=reactjs&paymentid=hegfheh"

func main() {

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
    fmt.Println(result.RawQuery)
	fmt.Println(result.Host)
     
	qparam := result.Query()

	for _,value := range qparam{
		fmt.Println("Value is : -",value)
	}
    
	//important &
    partsofurl := &url.URL{
		// Scheme: "https",
		// Host: "lco.dev",
		// Path: "/tutcss",
		
	}

      anotherurl := partsofurl.String()
	  fmt.Println(anotherurl)

}