package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int, error) { //error是固定写法error是个interface
	switch op {
	case "+":
		return a + b, nil
	case "/":
		q, _ := div(a, b) //不需要的返回值用下划线占位就可以
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported op:s%", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	//return a / b, a % b
	return q, r
}

func sumArgs(values ... int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func sumFunc(a, b int) int {
	return a + b
}
//接受函数为入参，函数编程 注意写法
func apply(opfun func(int, int) int, a, b int) int {
	p := reflect.ValueOf(opfun).Pointer()
	opName := runtime.FuncForPC(p).Name()
	result := opfun(a, b)
	fmt.Printf("calling func %s with args(%d,%d) ,get %d \n", opName, a, b, result)
	return result
}

func main() {
	q, r := div(6, 3)
	fmt.Println(q, r)
	fmt.Println(eval(2, 3, "+"))
	fmt.Println(eval(9, 3, "/"))

	if result, err := eval(5, 1, "/"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(sumArgs(1, 2, 3, 4, 5))

	apply(pow, 2, 3)
	apply(sumFunc, 2, 3)

	apply(func(i int, i2 int) int {
		return i * i2
	}, 15, 2)
}
