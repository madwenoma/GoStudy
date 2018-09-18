package main

import "fmt"

func sumInt(a, b int) int {
	return a + b
}

func someWontTestFunc() {
	s := []struct {
		name, age string
	}{
		{"tom", "11"},
		{"lee", "33"},
	}
	fmt.Println(s)
}
