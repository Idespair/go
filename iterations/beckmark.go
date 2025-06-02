package iteration

import "testing"

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration("a")
	}
}
