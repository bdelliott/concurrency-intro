package joke

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const dadJokeUrl = "https://icanhazdadjoke.com"


func jokeHandler(w http.ResponseWriter, r *http.Request) {

	req, err := http.NewRequest("GET", dadJokeUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	num, err := fmt.Fprintf(w, string(buf))
	if num != len(buf) {
		panic(err)
	} else if err != nil {
		panic(err)
	}
}

func Init() {
	// setup HTTP service
	http.HandleFunc("/", jokeHandler)
}

func Serve() {
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}