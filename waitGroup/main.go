package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usually these are pointers
func main() {

	websites := []string{
		"http://lco.dev",
		"http://google.com",
		"http://github.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("oops in endpoint")

	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}

}

/*
*Wait Group
what

why
*/
