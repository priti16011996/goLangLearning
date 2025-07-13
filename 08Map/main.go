package main

import "fmt"

func main() {
	fmt.Println("Go Lang Map ")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["Rb"] = "Rubby"
	fmt.Println("LanguageList: ", languages)
	delete(languages, "Rb")
	fmt.Println("LanguageList: ", languages)

	// loops are intresting

	for key, value := range languages {
		fmt.Printf("languages key %v and value %v\n ", key, value)
	}

}
