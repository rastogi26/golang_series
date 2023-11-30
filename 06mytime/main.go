package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome")
	Time := time.Now()
	fmt.Println(Time)

	fmt.Println(Time.Format("02-01-2006 Monday"))

}
