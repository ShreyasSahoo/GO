package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["TS"] = "typescript"
	languages["PY"] = "python"
	fmt.Println(languages)
	fmt.Println(languages["JS"])
	delete(languages, "TS")
	fmt.Println(languages)

	//iterating through a map
	for key, value := range languages {
		fmt.Println(key, value)
	}
}
