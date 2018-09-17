package main

import "fmt"

func fillSlice() {
	var arr []int
	for i := 0; i < 100; i++ {
		arr = append(arr, 2*i-1)
	}
	fmt.Println(arr)
}

func printArr(arr []int) {
	fmt.Printf("arr is %d,len=%d,cap=%d\n", arr, len(arr), cap(arr))
}

func main() {
	fillSlice()
	s1 := []int{1, 2, 4, 5}
	printArr(s1)
	s2 := make([]int, 16) //make的用法
	printArr(s2)
	s3 := make([]int, 16, 32)
	printArr(s3)

	fmt.Println("copy slice")
	copy(s2, s1)
	printArr(s2)

	fmt.Println("delete from slice")
	s2 = append(s2[:3], s2[4:]...) //删除中间元素 新语法
	printArr(s2)
	head := s2[0]
	fmt.Println(head)
	s2 = s2[1:]
	printArr(s2)
	//删除最后一个
	s2 = s2[:len(s2)-1]
	printArr(s2)

}
