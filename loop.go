package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

func conToBin(n int) string {
	binStr := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		binStr += strconv.Itoa(lsb)
	}
	return binStr
}

func readFile() {
	filename := "abc.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("forever...")
	}
}

func main() {
	fmt.Println(
		conToBin(2),
		conToBin(13),
		conToBin(2344),
		conToBin(1024),
	)
	readFile()
	//forever()
}
