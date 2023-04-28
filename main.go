package main

import (
	"log"
	"practice_go/sub1"
	"practice_go/sub2"
)

func main() {
	test1 := sub1.TestSub1(1)
	test1_1 := sub1.TestSub1_1(1)
	test2 := sub2.TestSub2(1)
	log.Printf("test1:%v\n", test1)
	log.Printf("test1_1:%v\n", test1_1)
	log.Printf("test2:%v\n", test2)
}
