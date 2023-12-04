<<<<<<< HEAD
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	//use walrus operator as we dont know what will come in the input
	welcome := "Welocme user"
    fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rating of Pizza:")

	//comma ok || comma error  ok
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
    
	//printf and println
	fmt.Printf("Type is: %T",input)




}

=======
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	//use walrus operator as we dont know what will come in the input
	welcome := "Welocme user"
    fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rating of Pizza:")

	//comma ok || comma error  ok
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
    
	//printf and println
	fmt.Printf("Type is: %T",input)




}

>>>>>>> ad16da4919a7c4062a63bfac2c8a7ca29dad1f9e
