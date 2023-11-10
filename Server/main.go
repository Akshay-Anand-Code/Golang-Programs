package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//PerformGetReq()
	//PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetReq() {
	const url = "http://localhost:8000/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	  {
		"course":"Let's go with golang",
		"price": 0,
		"site": "learnAkshay.in"

		
	  }
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/post"

	data := url.Values{}
	data.Add("First name", "Akshay")
	data.Add("Last name", "Anand")
	data.Add("UID", "2435")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
