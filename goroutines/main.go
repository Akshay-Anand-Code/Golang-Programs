package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usually these are pointers

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://google.com",
		"https://github.com",
		"https://facebook.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("OOPS in endpoint")
	}

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

}

/*
** GO-ROUTINES**

* Managed by Go runtime. Flexible stack - 2kb
* it helps to implement paralleism

 */
