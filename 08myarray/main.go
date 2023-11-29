package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	// var fruitList [4]string
	// fruitList[0] ="Apple"
	// fruitList[2]="Orange"
	// fruitList[3]="Peach"

	var fruitList =[4] string{"Apple","Orange","Peach"}

	fmt.Println("List is" ,fruitList)
	fmt.Println("Length is", len(fruitList))
}