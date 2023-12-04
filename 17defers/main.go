<<<<<<< HEAD
package main

import "fmt"

func main() {
	// 1 execute at last 
	// 2 LIFO
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")

	mydefer()

}

func mydefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
=======
package main

import "fmt"

func main() {
	// 1 execute at last 
	// 2 LIFO
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")

	mydefer()

}

func mydefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
>>>>>>> ad16da4919a7c4062a63bfac2c8a7ca29dad1f9e
}