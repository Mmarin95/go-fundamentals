package main

import (
	"fmt"
	"os"

	"miqgo.com/go/files/filemanagement"
)

func main() {
	rootDir, _ := os.Getwd()
	fmt.Println(rootDir)

	content, err := filemanagement.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(content)

	newContent := fmt.Sprintf("%s\n%s", content, "bbbb")

	err = filemanagement.WriteToFile("test2.txt", newContent)

	if err != nil {
		panic(err)
	}

}
