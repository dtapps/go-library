package system

import "testing"

func TestIsMac(t *testing.T) {
	t.Log(IsMac())
}

func BenchmarkIsMac(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsMac()
	}
}

func TestIsWindows(t *testing.T) {
	t.Log(IsWindows())
}

func BenchmarkIsWindows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsWindows()
	}
}

func TestIsLinux(t *testing.T) {
	t.Log(IsWindows())
}

func BenchmarkIsLinux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsWindows()
	}
}
