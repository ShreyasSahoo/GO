package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int32
	Platform string
	Password string   `json:"-"`              //hides the field
	Tags     []string `json:"tags,omitempty"` //omits empty field
}

func main() {
	// EncodeJson()
	decodeJson()
}

func EncodeJson() {
	webCourses := []course{
		{"React", 1000, "Udemy", "test@123", []string{"web-dev", "js"}},
		{"Nextjs", 1200, "Udemy", "test@342", []string{"front-end", "ts"}},
		{"Node", 700, "Coursera", "test@987", nil},
	}
	finalJson, err := json.Marshal(webCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)

	IndentedJson, err := json.MarshalIndent(webCourses, "", "\t") // prefix and indent - for json
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", IndentedJson)

}
func decodeJson() {
	jsonData := []byte(
		`[
			{
					"name": "React",
					"Price": 1000,
					"Platform": "Udemy",
					"tags": [
							"web-dev",
							"js"
					]
			},
			{
					"name": "Nextjs",
					"Price": 1200,
					"Platform": "Udemy",
					"tags": [
							"front-end",
							"ts"
					]
			},
			{
					"name": "Node",
					"Price": 700,
					"Platform": "Coursera"
			}
			]
	`,
	)
	var lcoCourse course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	//some cases are where you want to add data to key values
	// var myOnlineData map[string]interface{}
}
