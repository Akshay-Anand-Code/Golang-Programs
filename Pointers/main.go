package main

import "fmt"

type Person struct {
	name string
	age  int
}

func modifyAge(per *Person, age int) {
	per.age = age
}

func modify1(per *Person, name string) {
	per.name = name
}

func modifyInt(num *int, newValue int) {
	*num = newValue
}

func main() {
	person := Person{
		name: "Akshay",
		age:  20,
	}

	modifyAge(&person, 45)
	modify1(&person, "akash")
	fmt.Println(person)

	number := 42
	fmt.Println(number)
	modifyInt(&number, 8)
	fmt.Println(number)

}
