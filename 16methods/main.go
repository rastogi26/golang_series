package main

import "fmt"

func main() {

	// no inheritence , no parent or super or child in golang

	fmt.Println("Welcome")
	user1 := User{"dhruv","hello@gmail.com",true,21}
	fmt.Println(user1)
	fmt.Printf("User1 details are : %+v \n",user1)
	fmt.Printf("Name is %v ans email is %v \n",user1.Name,user1.Email)
	user1.GetStatus()
}

//important always use capital letter of first impiles public ie can export
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//method
func (u User) GetStatus(){
   fmt.Println("Is user active :- ",u.Status)
}