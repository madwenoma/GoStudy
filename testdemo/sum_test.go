package main

import "testing"

/**
go的测试风格是表格测试，table，开头slice struct定义一堆数据
https://www.bilibili.com/video/av24365381/?p=33
go test：
1.test文件结尾必须是_test， 引入import "testing"
2.文件名必须以Test开头定义函数
3.t *testing.T 是入参
4.t.Errorf
5.命令行也可以运行 go test .
6.命令行查看junit代码覆盖率
	6.1.go test -coverprofile=c.out //该句运行test并生成结果c.tou
	6.2.go tool cover -html=c.out   //该句转为html并在默认浏览器中打开，绿色为覆盖，红色为未覆盖的部分
 */

func TestSum(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 7},
		{5, 5, 10},
		{22, 33, 55},
	}

	for _, tt := range tests {
		if actual := sumInt(tt.a, tt.b); actual != tt.c {
			t.Errorf("sum(%d, %d) get %d,expect %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

//性能测试
//go test -bench .
func BenchmarkSum(t *testing.B) {
	a, b := 22, 33
	c := 55

	for i := 0; i < t.N; i ++ {
		if actual := sumInt(a, b); actual != c {
			t.Errorf("sum(%d, %d) get %d,expect %d", a, b, actual, c)
		}
	}
}
