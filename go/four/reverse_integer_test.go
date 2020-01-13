package four

import (
	"testing"
)

func TestReverse(t *testing.T) {
	i := 123
	r := reverse(i)
	if r != 321 {
		t.Errorf("reverse(123) failed. Got %d, expected 321.", r)
	}

	i = 1534236469
	r = reverse(i)
	if r != 0 {
		t.Errorf("reverse(1534236469) failed. Got %d, expected 0.", r)
	}
}
