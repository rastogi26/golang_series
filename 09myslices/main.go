<<<<<<< HEAD
package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple","Orange","Peach"}

	fmt.Println("List is ",fruitList)
	fruitList = append(fruitList, "Mango","Banana")
	fmt.Println("Updated List ",fruitList)
     
    fmt.Println(fruitList[2:4])
	fmt.Printf("Type is %T \n",fruitList)
	fmt.Println("Length is ", len(fruitList))
    
	//2nd method 
	highScores :=  make([]int,4)
	highScores[0]=10
	highScores[1]=15
	highScores[2]=98
	highScores[3]=-7

	highScores = append(highScores,1,3,2,5 )
	fmt.Println("Scores is", highScores)
	fmt.Println("Scores Length is",len(highScores))
    fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("High Scores Sorted in Ascending Order:",highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	//how to remove value from slice
	var courses = []string{"reactjs","python","ruby","swift","c++"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]... )
	fmt.Println("Updated Courses:", courses)




=======
package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple","Orange","Peach"}

	fmt.Println("List is ",fruitList)
	fruitList = append(fruitList, "Mango","Banana")
	fmt.Println("Updated List ",fruitList)
     
    fmt.Println(fruitList[2:4])
	fmt.Printf("Type is %T \n",fruitList)
	fmt.Println("Length is ", len(fruitList))
    
	//2nd method 
	highScores :=  make([]int,4)
	highScores[0]=10
	highScores[1]=15
	highScores[2]=98
	highScores[3]=-7

	highScores = append(highScores,1,3,2,5 )
	fmt.Println("Scores is", highScores)
	fmt.Println("Scores Length is",len(highScores))
    fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("High Scores Sorted in Ascending Order:",highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	//how to remove value from slice
	var courses = []string{"reactjs","python","ruby","swift","c++"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]... )
	fmt.Println("Updated Courses:", courses)




>>>>>>> ad16da4919a7c4062a63bfac2c8a7ca29dad1f9e
}