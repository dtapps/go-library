package gouuid

import (
	"testing"
)

func TestGetUuId(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{{}, {}, {}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(GetUuId())
		})
	}
}

func BenchmarkGetUuId(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetUuId()
	}
}
