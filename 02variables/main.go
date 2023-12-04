<<<<<<< HEAD
package main

import "fmt"
//public
const Login string = "hebaiyabvi" 
//private 
// const registeration string = "hdgfhgf"
func main() {
    //  var username complex128
	//  fmt.Println(username)
	//  fmt.Printf("username type is : %T \n",username)

	//  var  isusername bool  = true
	//  fmt.Println(isusername)
	//  fmt.Printf("username type is : %T \n ",isusername)
	 
	//  fmt.Println(Login)

	  var name string
	  fmt.Println("string is",name)
	  var number int
	  fmt.Println("number is ", number)
	 
	  //implicit type
	  var website = "learn.com"
	  fmt.Println(website)
	  // it automatically selects the type based on the initalise
	  // if we do like below it gives error because it is string not int
	//   website = 3

	// no  var style (walrus opator)
	numberOfUser :=90000
	fmt.Println(numberOfUser)

=======
package main

import "fmt"
//public
const Login string = "hebaiyabvi" 
//private 
// const registeration string = "hdgfhgf"
func main() {
     var username complex128
	 fmt.Println(username)
	 fmt.Printf("username type is : %T \n",username)

	 var  isusername bool  = true
	 fmt.Println(isusername)
	 fmt.Printf("username type is : %T \n ",isusername)
	 
	 fmt.Println(Login)
	 
>>>>>>> ad16da4919a7c4062a63bfac2c8a7ca29dad1f9e
}