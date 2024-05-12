package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		ch1 chan byte = make(chan byte, 1)
	)

	fmt.Println(len(ch1))
	ch1 <- 0x00
	fmt.Println(len(ch1))
	time.Sleep(3 * time.Second)
	res := <-ch1
	fmt.Println("result:", res)
}
