package main

import (
	"fmt"
	"io/ioutil"
)

func ifDemo() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename) //返回两个参数byte[]和err
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents) //%s可以打印数组
	}

	//下面写法更简单 但要注意if后面的语句定义的变量作用于就是这if块里
	if txtContent, errMsg := ioutil.ReadFile(filename); errMsg != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", txtContent)
	}
}

//switch demo 可以不需要表达式
func grade(score int) string {
	result := "";
	switch {
	case score < 60:
		result = "F"
	case score < 80:
		result = "D"
	case score < 90:
		result = "B"
	case score <= 100:
		result = "A"
	default:
		panic("wrong score")
	}
	return result;
}
func main() {
	ifDemo();
	fmt.Println(grade(55), grade(70), grade(89), grade(99), grade(100))
}
