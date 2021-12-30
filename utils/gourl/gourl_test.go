package gourl

import (
	"testing"
	"time"
)

func TestLenCode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "https://www.dtapp.net",
	}, {
		name: "https://www.dtapp.net",
	}, {
		name: "https://www.dtapp.net",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			t.Log(LenCode(tt.name))
			elapsed := time.Since(start)
			t.Log("run time", elapsed)
		})
	}
}

func BenchmarkLenCode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LenCode("https://www.dtapp.net")
	}
}

func TestDeCode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "https%3A%2F%2Fwww.dtapp.net",
	}, {
		name: "https%3A%2F%2Fwww.dtapp.net",
	}, {
		name: "https%3A%2F%2Fwww.dtapp.net",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			t.Log(DeCode(tt.name))
			elapsed := time.Since(start)
			t.Log("run time", elapsed)
		})
	}
}

func BenchmarkDeCode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DeCode("https%3A%2F%2Fwww.dtapp.net")
	}
}
