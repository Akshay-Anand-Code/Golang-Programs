package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 232
	highScores[3] = 543
	highScores = append(highScores, 222, 111, 333)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	var books = []string{"maths", "science", "history", "english", "geometry", "linguistic"}
	fmt.Println(books)
	var index int = 3
	books = append(books[:index], books[index+1:]...)
	fmt.Println(books)

}
