package main

import "fmt"

//星号代表指针操作
func swap(a, b *int) {
	*a, *b = *b, *a
}

func swapOrigin(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 2, 3
	swap(&a, &b) //调用的时候用&取得指针
	fmt.Println(a, b)
	fmt.Println(swapOrigin(3, 4))
}
