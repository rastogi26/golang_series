<<<<<<< HEAD
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Rate the pizza:")

	reader:=bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating :",input)
    
	numRating,err:= strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	}else
	{
		fmt.Println("Adding 1 to your rating:",numRating+1)
	}

=======
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Rate the pizza:")

	reader:=bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating :",input)
    
	numRating,err:= strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	}else
	{
		fmt.Println("Adding 1 to your rating:",numRating+1)
	}

>>>>>>> ad16da4919a7c4062a63bfac2c8a7ca29dad1f9e
}