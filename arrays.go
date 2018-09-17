package main

import "fmt"

func printlnArr(arr [3]int) {
	for key, value := range arr {
		fmt.Println(key, value)
	}
}

func main() {
	arr := [3]int{1, 2, 3}
	var arr2 [5]int
	arr3 := [...]int{23, 4}
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)

	var grid [4][5]bool
	fmt.Println(grid)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//for _, v := range arr3 //可以直接不用下标 i
	for i, v := range arr3 { //range的用法，可以直接获得值 v直接使用
		fmt.Println(arr3[i])
		fmt.Println(v)
	}

	printlnArr(arr)
	//printArr(arr2)//报错 arr2 长度不是3

}
