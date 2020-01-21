package five

import "testing"

func TestMyAtoi(t *testing.T) {
	var i string
	var r int

	i = "42"
	r = myAtoi(i)
	if r != 42 {
		t.Errorf("myAtoi(42) failed. Got %d, expected 42.", r)
	}

	i = "011"
	r = myAtoi(i)
	if r != 11 {
		t.Errorf("myAtoi(011) failed. Got %d, expected 11.", r)
	}

	i = "   01134dfgg   "
	r = myAtoi(i)
	if r != 1134 {
		t.Errorf("myAtoi(01134dfgg) failed. Got %d, expected 1134.", r)
	}

	i = "d134dfgg   "
	r = myAtoi(i)
	if r != 0 {
		t.Errorf("myAtoi(d134dfgg) failed. Got %d, expected 0.", r)
	}
}
