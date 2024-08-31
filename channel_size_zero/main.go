package main

import (
	"fmt"
	"time"
)

func main() {
	const MAX_NUM_GOROUTINES int = 100
	ch := make(chan byte, 1)
	startRead := make(chan bool)

	for i := 0; i < MAX_NUM_GOROUTINES; i++ {
		go func(c chan byte, idx int) {
			fmt.Printf("goroutine %v start\n", idx)
			time.Sleep(time.Duration(idx*100) * time.Millisecond)
			fmt.Printf("goroutine %v c <- %v\n", idx, idx)
			if idx == MAX_NUM_GOROUTINES-1 {
				startRead <- true
			}
			c <- byte(idx)
			fmt.Printf("goroutine %v end\n", idx)
		}(ch, i)
	}

	<-startRead
	fmt.Printf("len(ch):%v\n", len(ch))

	for b := range ch {
		fmt.Printf("%v <- ch\n", b)
		if len(ch) == 0 {
			break
		}
	}

	fmt.Println("fin")
}
