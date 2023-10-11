package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://localhost:8000"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength) //-1 when length is unknown
	fmt.Println(content)

	//another way to write get request
	var responseString strings.Builder
	byteContent, _ := responseString.Write(content)

	fmt.Println(byteContent)
	fmt.Println(responseString.String())

}
