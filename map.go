package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m1 := make(map[string]int) //empty map
	var m2 map[int]string      //nil
	fmt.Println(m, m1, m2)

	//遍历
	for key, value := range m {
		fmt.Println(key, value) //每次输出顺序不是固定的
	}

	//
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	causeName, ok := m["cause"] //为什么ok可以重复，因为前面的causeName是新定义的，如果两个都是旧的变量名，就会报错
	fmt.Println(causeName, ok)

	//
	fmt.Println("delete values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
