package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("smeeee")
	content := "This needs to go in a file - EduBytes"

	file, err := os.Create("./mylcogofile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is\n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
