package deferdemo

import (
	"io"
	"fmt"
	"strings"
)

func Fibonacci() FiboFunc {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type FiboFunc func() int

func (p FiboFunc) Read(b []byte) (n int, err error) {
	next := p() //调用一下就是下一个
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(b)
}