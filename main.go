package main

import (
	"fmt"
	"time"
)

func processA(c chan<- string) {
	time.Sleep(1 * time.Second)
	c <- "result a"
}

func processB(c chan<- string) {
	time.Sleep(1 * time.Second)
	c <- "result b"
}

func processC(c chan<- string) {
	time.Sleep(1 * time.Second)
	c <- "result c"
}

func processD(c chan<- string) {
	time.Sleep(1 * time.Second)
	c <- "result d"
}

func main() {
	start := time.Now()

	c := make(chan string, 4)

	go processA(c)
	go processB(c)
	go processC(c)
	go processD(c)

	fmt.Printf("res1: %v\n", <-c)
	fmt.Printf("res2: %v\n", <-c)
	fmt.Printf("res3: %v\n", <-c)
	fmt.Printf("res4: %v\n", <-c)

	fmt.Printf("Use time : %v\n", time.Since(start))
}
