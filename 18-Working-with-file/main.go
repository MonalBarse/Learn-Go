package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Files in Golang")

	content := "This is a test file"

	file, err := os.Create("./test.txt")

	if err != nil {

		fmt.Println(err)
		panic(err)
	}
	length, err := io.WriteString(file, content) // This will write the content to the file
	// length will return the length of the content written to the file
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Length: ", length)
	defer file.Close() // This will close the file after the function completes
	readFile("./test.txt")

}

func readFile(filename string) {
	// Read the file and print the content
	/* 	file, err := os.Open(filename)
	  checkNilError(err)

		defer file.Close()
		fileinfo, err := file.Stat()
	  checkNilError(err)

		filecontent := make([]byte, fileinfo.Size())

		_, err = file.Read(filecontent)
	  checkNilError(err)
		fmt.Println("Content of the file: ", string(filecontent)) */

	filecontent, err := os.ReadFile(filename)
	checkNilError(err)

	// Print the content of the file
	fmt.Println("Content of the file: ", string(filecontent))
}

func checkNilError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
