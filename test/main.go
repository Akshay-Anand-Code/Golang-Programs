package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) changeName(nameChange string) {
	p.name = nameChange
}

func (p *Person) mutateName(nameChange string) {
	p.name = nameChange
}

func main() {
	akshay := &Person{
		name: "akshay",
		age:  20,
	}

	fmt.Println(akshay)
	akshay.changeName("abhi")
	fmt.Println(akshay)
	akshay.mutateName("abhi")
	fmt.Println(akshay)
}
