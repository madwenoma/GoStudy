package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	s := "yes我爱慕课网!"

	fmt.Println(len(s))
	for _, value := range []byte(s) { //utf8
		fmt.Printf("%X", value)
	}
	fmt.Println()

	for key, value := range s { //unicode
		fmt.Printf("%d %X\n", key, value)
	}

	fmt.Printf("count:%d\n", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //utf8字节转rune
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}

	fmt.Println()
	for key, value := range []rune(s) {
		fmt.Printf("%d %c\n", key, value)
	}

	strs := strings.Split(s, "爱")
	fmt.Println(strs[0])
	fmt.Println(strs[1])
	strs = strings.Fields(s)
	fmt.Println(strs)
	fmt.Println(strings.Contains(s,"!"))
	fmt.Println(strings.Index(s,"!"))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Join([]string{"apple","pear"},"/"))
}
