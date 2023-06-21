package main

import "fmt"

func main() {
	fmt.Println("Learning Maps on go")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	languages["go"] = "golang"
	languages["ts"] = "typescript"

	fmt.Println("List of all languages", languages)
	fmt.Println("Length of languages", len(languages))
	fmt.Println("Language for py is", languages["py"])

	delete(languages, "rb")
	fmt.Println("List of all languages", languages)

	for key, value := range languages {
		fmt.Printf("%s => %s\n", key, value)
	}
}
