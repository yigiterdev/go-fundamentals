package main

import (
	"fmt"
	"os"

	fileutils "yigiterdev/files/file-management"
)

func main() {
	dir, _ := os.Getwd()
	newContent := fmt.Sprintf("Hello, World! %s", dir)
	fileutils.WriteToFile(dir+"/file-management/data.txt", newContent)
	content, _ := fileutils.ReadFile(dir + "/file-management/data.txt")
	fmt.Println(content)
}
