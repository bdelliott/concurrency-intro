package main

import (
	"fmt"
	"github.com/bdelliott/concurrency-intro/go/joke"
)


func main() {
	// parallelize joke fetching client
	jokes := joke.Fetch()

	fmt.Println("I've got jokes:")
	for _, joke := range jokes {
		fmt.Println("Joke: ")
		fmt.Println(joke.Joke)
		fmt.Println("-------------------")
	}
}