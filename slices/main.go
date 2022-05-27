package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Slicesss")

	var fruitlist = []string{"Apple", "Tomato"}
	// fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Peach", "Kaku")
	// fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	// fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0] = 2345
	highScores[1] = 9456
	highScores[2] = 2347
	highScores[3] = 2348
	// highScores[4] = 7777

	highScores = append(highScores, 666)

	// fmt.Println(highScores)

	sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove value from slice based on index

	var courses = []string{"Javascript", "GoLang", "Python", "Ruby"}

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	

}
