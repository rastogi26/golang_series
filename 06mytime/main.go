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
	createDate:= time.Date(2020,time.March,10,23,23,0,0,time.UTC)
	fmt.Println(createDate)

}
