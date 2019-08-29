package main

import (
	"fmt"
	"time"
)

// ticker is for repeat at regular interval of time
func main() {
	ticker :=time.NewTicker(1000* time.Millisecond)
	//ticker.C is of channel type
	for t := range ticker.C{
		fmt.Println("Tick at", t)
	}
	// it also have stop - unless if it is called, it will not be available for the garbage collection
	//newticketer has stop and ticker dont have stop
	// to stop the ticker we have to use the ticker.stop
	//time.Sleep(1600 * time.Millisecond)
	// STOP
	//ticker.Stop()
}
