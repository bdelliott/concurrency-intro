package main

import "github.com/bdelliott/concurrency-intro/go/joke"

func main() {
	joke.Init()
	joke.Serve()
}
