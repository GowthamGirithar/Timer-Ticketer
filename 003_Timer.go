package main

import (
	"fmt"
	"time"
)

func main() {
	//it will not repeat at regular interval and it is only one time
	//it also returns of type channel
	// it also have stop - unless if it is called, it will not be available for the garbage collection
	// so newtimer is best to use than timer
	timer := time.NewTimer(100 *time.Millisecond)
	go func() {
		fmt.Println(<- timer.C)
	}()

	time.Sleep(10000 *time.Millisecond)
	
}
