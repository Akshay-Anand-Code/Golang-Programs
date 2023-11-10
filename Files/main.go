package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Files in golang")
	content := "This is the content of the file"

	file, err := os.Create("./myFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myFile.txt")

}

func readFile(filename string) {
	ioutil.ReadFile(filename)
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println(databyte)
	fmt.Println(string(databyte))
}
