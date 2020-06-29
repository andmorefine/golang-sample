package integers

import "testing"

func TestAverage(t *testing.T) {

	got := Average(10, 20)
	want := 15

	if got != want {
		t.Errorf("average got '%d' wont '%d'", got, want)
	}
}
