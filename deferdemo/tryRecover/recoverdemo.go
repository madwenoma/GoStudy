package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()

		if err, ok := r.(error); ok { //如果转成了error，就是正常的error，否则就不是error类型
			fmt.Println("error:", err)
		} else {
			panic(fmt.Sprintf("unkonw error %v", r))
		}
	}()
	//panic(123)
	a, b := 5, 0
	fmt.Println(a / b)
}

func main() {
	tryRecover()
}
