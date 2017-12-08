package lineproto

import (
	"io/ioutil"
	"testing"
)

func Benchmark_NewParser(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = newParser("lineproto", []byte("cpu value=42\n"))
	}
}

func Benchmark_Minimal(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := Parse("lineproto", []byte("cpu value=42\n"))
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_500lines(b *testing.B) {
	metrics, err := ioutil.ReadFile("500lines")
	if err != nil {
		panic("error reading file")
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := Parse("lineproto", metrics)
		if err != nil {
			panic(err)
		}
	}
}
