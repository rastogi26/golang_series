package main

import "fmt"

func main() {
	greeting()
	result:= add(3,4)
	fmt.Println("Result is", result)

     proRes,message := proAdder(2,4,3,4,5,1,2,3,34)
	 fmt.Println("Pro_Result is ",proRes)
	 fmt.Println(message)
}

//pro functions  ... are variatic functions and can expect various value
// also can return two datatypes 
func proAdder(values... int)(int,string){
	total :=0
	
	for _,val := range values{
		total += val;
	}
	return total,"Job Done"
}

// function signatures 
func add(val1 int , val2 int) int{
	return (val1+val2)
}

func greeting() {
	fmt.Println("Namaste")
}