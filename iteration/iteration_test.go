package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("fine")
	want := "finefinefine"

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("fine")
	}
}
