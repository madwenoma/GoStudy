package main

//P177

//通道本身就是一个并发安全的队列 ，适合用来作为ID生成器、pool等场景
type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var b []byte
	select {
	case b = <-p:
	default:
		b = make([]byte, 10)
	}
	return b
}

func (p pool) put(v []byte) {
	select {
	case p <- v:
	default:
	}
}
