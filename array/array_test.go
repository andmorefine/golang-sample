package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}

	t.Run("ここで失敗", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 14

		if got != want {
			t.Errorf("got %v want %d given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 3, 4})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected got %v but want %v", got, want)
	}

}
