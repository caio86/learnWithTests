package main

import "testing"

func BenchmarkWalk(b *testing.B) {
	structs := make([]any, 100)
	for i := range structs {
		structs[i] = Profile{1, "test"}
	}

	b.ResetTimer()

	for range b.N {
		walk(structs, func(s string) {})
	}
}
