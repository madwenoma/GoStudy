package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7,}
	s := arr[2:6]
	fmt.Println(s)
	fmt.Println(arr[:6])
	s1 := arr[2:]
	fmt.Println(s1)
	s2 := arr[:] //取全部
	fmt.Println(s2)

	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	//re slice
	fmt.Println("Reslice...")
	s2 = s2[:4]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	//extending slice
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = nums[2:6] //虽然取到5，但6和7还存在，因为silice机制，向后有cap
	fmt.Printf("lens:%d,cap:%d,s1=%d\n", len(s1), cap(s1), s1)
	s2 = s1[3:5]
	fmt.Printf("lens:%d,cap:%d,s2=%d\n", len(s2), cap(s2), s2)

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt.Println(nums)

}

func updateSlice(arr []int) {
	arr[0] = 100
}
