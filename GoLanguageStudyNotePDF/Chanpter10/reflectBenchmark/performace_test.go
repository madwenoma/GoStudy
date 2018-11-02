package main

import "testing"

/**

BenchmarkSet-4           2000000000               	0.68 ns/op            	0 B/op          0 allocs/op
BenchmarkRSet-4          5000000                 	256 ns/op               8 B/op          2 allocs/op
BenchmarkCRSet-4         20000000                	66.1 ns/op             	4 B/op          1 allocs/op

BenchmarkCall-4          1000000000               	2.69 ns/op            	0 B/op          0 allocs/op
BenchmarkRCall-4         5000000                  	323 ns/op               8 B/op          1 allocs/op


 */
func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set(100)
	}
}

func BenchmarkRSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rset(100)
	}
}

func BenchmarkCRSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cachedrset(100)
	}
}


func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		directCall()
	}
}

func BenchmarkRCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rCall()
	}
}