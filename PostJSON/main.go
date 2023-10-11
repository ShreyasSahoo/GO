package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformPostJSONRequest()
	performPostFormRequest()
}

func PerformPostJSONRequest() {
	const myURL = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
	{
		"coursename":"golang",
		"price":"5000"
	}
	`)
	response, err := http.Post(myURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func performPostFormRequest() {
	const myURL = "http://localhost:8000/postForm"

	data := url.Values{}
	data.Add("firstName", "Shreyas")
	data.Add("lastName", "Sahoo")
	data.Add("email", "shreyassahoo26@gmail.com")

	response, err := http.PostForm(myURL, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}
