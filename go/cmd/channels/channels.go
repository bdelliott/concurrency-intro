package main

import (
"fmt"
"github.com/bdelliott/concurrency-intro/go/adder"
"time"
)


func main() {

	counter := 0
	n := 1000000

	start := time.Now()
	adder.ChanneledAdder(&counter, n)
	elapsed := time.Since(start)

	fmt.Printf("counter = %d\n", counter)
	fmt.Println("Wall-clock time: ", elapsed)

}