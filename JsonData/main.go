package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int
	Platform string
	Password string `json:"-"`
	Tags     []string
}

func main() {
	//encodeJson()
	DecodeJson()
}

func encodeJson() {

	courses := []course{
		{"WebDev", 200, "akshay.com", "aks123", []string{"webdev", "learnwevdev"}},
		{"c++", 300, "lco.com", "ke0123", []string{"c++ learn", "coding"}},
		{"java", 5500, "java.com", "cii123", []string{"java learn", "coding"}},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	
	{
		"name": "java",
		"Price": 5500,
		"Platform": "java.com",
		"Tags": [
				"java learn",
				"coding"
		]
    }
           
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json data was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
