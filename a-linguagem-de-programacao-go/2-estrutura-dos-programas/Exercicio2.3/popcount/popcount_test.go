package popcount

import (
	"testing"
)

var testvar uint64

func BenchmarkPopCount(b *testing.B) {
	testvar = 64
	for n := 0; n < b.N; n++ {
		PopCount(testvar)
	}

}

func BenchmarkPopCount2(b *testing.B) {
	testvar = 64
	for n := 0; n < b.N; n++ {
		PopCount2(testvar)
	}

}
