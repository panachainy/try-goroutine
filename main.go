package main

import (
	"fmt"
	"time"
)

func processA() {
	time.Sleep(1 * time.Second)
	fmt.Println("result a")
}

func processB() {
	time.Sleep(1 * time.Second)
	fmt.Println("result b")
}

func processC() {
	time.Sleep(1 * time.Second)
	fmt.Println("result c")
}

func processD() {
	time.Sleep(1 * time.Second)
	fmt.Println("result d")
}

func main() {
	start := time.Now()
	go processA()
	go processB()
	go processC()
	go processD()
	fmt.Printf("Use time : %v\n", time.Since(start))
}
