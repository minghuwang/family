package main

import "testing"

func BenchmarkFib10(b *testing.B) {
	for n := 0; n<b.N; n++ {
		Fib(10)
	}
}
func BenchmarkFib100(b *testing.B) {
	for n := 0; n<b.N; n++ {
		Fib(100)
	}
}

