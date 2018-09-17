package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	file, err := os.Open("abc.txt")
	if err != nil {
		panic("file not exist")
	}
	printFileContents(file)
	//read string

	str := `hehe , \ "" is some 
			txt
			exe`
	printFileContents(strings.NewReader(str))

}
