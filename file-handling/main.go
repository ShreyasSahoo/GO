package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("file handling in Golang")
	content := "This will be the content of the file"

	file, err := os.Create("./new-file.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("the length is ", length)
	defer file.Close()

}
func ReadFile(filename string) {
	dataBuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataBuffer))
}
