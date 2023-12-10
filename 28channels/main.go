package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang-  LearnCodeOnline.in")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{} 

	// fmt.Println(<-myCh)
	// push  into channel
	// myCh <- 5
	wg.Add(2)
	// Recevie ONLY  <-chan
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		//fmt.Println(<-myCh)  // listener

		wg.Done()
	}(myCh, wg)
	// send ONLY chan<-
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}