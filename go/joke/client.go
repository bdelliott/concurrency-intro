package joke

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

type Joke struct {
	Id	string
	Joke string
	Status int
}

const proxyUrl = "http://localhost:9000"


func Fetch() []Joke {
	numWorkers := runtime.NumCPU()
	fmt.Printf("Splitting work across %d workers.\n", numWorkers)

	jokes := make([]Joke, numWorkers)
	c := make(chan Joke)

	for i := 0; i < numWorkers; i++ {
		go fetchJoke(c)
		joke := <- c
		jokes[i] = joke
	}
	return jokes
}


func fetchJoke(c chan Joke) {
	resp, err := http.Get(proxyUrl)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(body)
	}

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		panic(err)
	}
	c <- joke
}