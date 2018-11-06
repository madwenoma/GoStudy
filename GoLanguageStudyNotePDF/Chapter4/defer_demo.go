package main

import "fmt"

func main() {
	//x, y := 1, 2
	//defer func(a int) {
	//	fmt.Println("defer x,y=", a, y)
	//}(x)
	//x += 100
	//y += 200
	//fmt.Println(x, y)
	fmt.Println("==============")
	fmt.Println("test:", testDefer())

}

func testDefer() (z int) {
	defer func() {
		fmt.Println("defer:", z)
		z += 100 //修改返回值
	}()
	return 100 //实际执行次序 z=100 ,call defer,ret
}
