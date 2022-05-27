package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is: ", result)
	proResult, myMesssage := proAdder(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(proResult, myMesssage)

}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Heloo pro"
}

func greeter() {
	fmt.Println("Asalamo alaikum")
}
