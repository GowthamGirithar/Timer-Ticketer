package main

import (
	"fmt"
	"time"
)

func main() {
	//Tick will not have the stop to terminate the process
	ticker := time.Tick(100*time.Millisecond)
	for t :=range  ticker{
		fmt.Println(t)
	}

}
