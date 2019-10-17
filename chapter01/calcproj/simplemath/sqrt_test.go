package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	t.Log("test sqrt")
	{
		v := Sqrt(64)
		if v != 8 {
			t.Fatal("err", v)
		}
	}
}
