package main

import (
	"fmt"

)

func main() {

	language := make(map[string]string)
	
	language["js"] ="javascript"
	language["rb"] ="ruby"
	language["py"] ="python"

	fmt.Println(language)
    //delete
	delete(language,"rb")
	fmt.Println(language)

	//loop is interesting

	for key,value := range language{
		fmt.Printf("For key %v , value is %v \n",key,value)
	}



}