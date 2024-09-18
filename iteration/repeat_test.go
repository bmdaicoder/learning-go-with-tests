package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeatCount := 10
	repeated := Repeat("a", repeatCount)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 300000)
	}
}
