<<<<<<< HEAD
package main

import "fmt"

func main() {
	// days := []string{"Sunday", "Monday", "Tuesday", "Wenesday", "Thusday", "Friday", "Saturday"}
    
	//method1
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//method2
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	//method3
	// for index,day := range days{
	// 	fmt.Printf("Index %v has %v \n",index,day)
	// }

	// no ++_ in go 
	//method 4 while

	//break and continue are normal  , goto 
	count := 1
	
	for count<10{
		if count==2 {
			goto india
		}
		fmt.Println(count)
		count++
	}
	
	india:
		fmt.Println("I am in India ")
=======
package main

import "fmt"

func main() {
	// days := []string{"Sunday", "Monday", "Tuesday", "Wenesday", "Thusday", "Friday", "Saturday"}
    
	//method1
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//method2
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	//method3
	// for index,day := range days{
	// 	fmt.Printf("Index %v has %v \n",index,day)
	// }

	// no ++_ in go 
	//method 4 while

	//break and continue are normal  , goto 
	count := 1
	
	for count<10{
		if count==2 {
			goto india
		}
		fmt.Println(count)
		count++
	}
	
	india:
		fmt.Println("I am in India ")
>>>>>>> ad16da4919a7c4062a63bfac2c8a7ca29dad1f9e
}