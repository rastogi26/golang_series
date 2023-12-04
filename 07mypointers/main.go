<<<<<<< HEAD
package main

import "fmt"
func main()  {

	// why pointers exist
	// so when we pass variables to functions a copy is created \
	//but if we want to change the actural place then we use pointers ie use their memory address location
	fmt.Println("Pointers")

	var ptr1 *int
	fmt.Println("Value of pointer", ptr1)
	//therefore the value of pointer by default = nil

	myNumber := 23

	ptr2 := &myNumber
	fmt.Println("Adress location is ", ptr2)
	fmt.Println("Value is ", *ptr2)

=======
package main

import "fmt"
func main()  {

	// why pointers exist
	// so when we pass variables to functions a copy is created \
	//but if we want to change the actural place then we use pointers ie use their memory address location
	fmt.Println("Pointers")

	var ptr1 *int
	fmt.Println("Value of pointer", ptr1)
	//therefore the value of pointer by default = nil

	myNumber := 23

	ptr2 := &myNumber
	fmt.Println("Adress location is ", ptr2)
	fmt.Println("Value is ", *ptr2)

>>>>>>> ad16da4919a7c4062a63bfac2c8a7ca29dad1f9e
}