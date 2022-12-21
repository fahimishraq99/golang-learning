package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang-  LearnCodeOnline.in")

	myCh := make(chan int, 2) //creating channel
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	// Receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//<-chan int; meaning something is going outside the box. Receive only. Not even allowed to put a close() statement here otherwise an error will be thrown.
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		//line 20 to 23 are executed to check whether we are getting the true value or false value according to [close(mych)] command. If we don't use these lines then we will get same output (0) for both [1: mych <- 0 2: close (mych)] and [close(mych)] but will not understand which output (0) is true and which is false as well. Thus, to understand we should use these commands of line 20 to 23

		//fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) { //chan<- int; meaning something is going inside the box
		myCh <- 0
		close(myCh) //putting close(mych) before passing a value will cause return of an error. except you are using commands of line 20 to 23
		//myCh <- 5
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
