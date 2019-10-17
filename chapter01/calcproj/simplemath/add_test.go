package simplemath

import "testing"

func TestAdd(t *testing.T) {
	t.Log("test add")
	{
		v := Add(1, 2)
		if v != 3 {
			t.Fatal("err", v)
		}
	}
}
