package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello my name is Dhruv"

	file, err := os.Create("./file1.txt")

    checkNilErr(err)

	length ,err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is", length)
	defer file.Close()  
    readFile("./file1.txt")

}

func checkNilErr(err error){
  if err!=nil{
		panic(err)
	}

}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
     
	fmt.Println("The content inside file1 is \n",string(databyte))

}