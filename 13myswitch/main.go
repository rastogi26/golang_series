<<<<<<< HEAD
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6)+1

	fmt.Println("Dice number is ", diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("Move 1")
	case 2:
		fmt.Println("Move 2")	
	case 3:
		fmt.Println("Move 3")	
	case 4:
		fmt.Println("Move 4")
		fallthrough	
	case 5:
		fmt.Println("Move 5")	
	case 6:
		fmt.Println("Move 6")		
	default:
		fmt.Println("wooh!!")			
	}

	

=======
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6)+1

	fmt.Println("Dice number is ", diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("Move 1")
	case 2:
		fmt.Println("Move 2")	
	case 3:
		fmt.Println("Move 3")	
	case 4:
		fmt.Println("Move 4")
		fallthrough	
	case 5:
		fmt.Println("Move 5")	
	case 6:
		fmt.Println("Move 6")		
	default:
		fmt.Println("wooh!!")			
	}

	

>>>>>>> ad16da4919a7c4062a63bfac2c8a7ca29dad1f9e
}