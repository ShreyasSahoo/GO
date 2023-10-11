package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hgk87685"

func main() {
	fmt.Println("Welcome to handling URL")
	fmt.Println(myURL)

	//parsing the url
	result, _ := url.Parse(myURL)
	fmt.Println("Scheme ", result.Scheme)
	fmt.Println("HOST ", result.Host)
	fmt.Println("Path", result.Path)
	fmt.Println("Port", result.Port())
	fmt.Println("raw query", result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The Type of qParams is %T\n", qParams)

	for key, value := range qParams {
		fmt.Println(key, " - ", value)
	}
	partsOfURL := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",

		Path:    "learn",
		RawPath: "user=shreyas",
	}
	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
