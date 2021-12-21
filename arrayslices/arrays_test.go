package arrayslices_test

import (
	"reflect"
	"testing"

	"github.com/rssharma75/learn_go/arrayslices"
)

func assert(got, want int, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Test Failed, expected:%d, got:%v", want, got)
	}

}

func assertSlices(got, want []int, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Test Failed, expected:%d, got:%v", want, got)
	}
}
func TestArraySlices(t *testing.T) {
	t.Run("Test Sum", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		want := 15
		got := arrayslices.Sum(arr)

		assert(got, want, t)
	})

	t.Run("Variable number of slice inputs", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{5, 8}

		want := []int{3, 13}
		got := arrayslices.SumAll(a, b)

		assertSlices(got, want, t)
	})
}
