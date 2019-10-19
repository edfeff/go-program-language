package library

import "testing"

func TestMusicManagerOP(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Fatal("mm is nil")
	}
	if mm.Len() != 0 {
		t.Error("mm is not empty")
	}
	m0 := &MusicEntry{
		Id:     "1",
		Name:   "A",
		Artist: "AA",
		Source: "AAA",
		Type:   "AAAA",
	}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("len is not 1")
	}
	m := mm.Find("A")
	if m == nil {
		t.Error("find is fail")
	}
	if m.Id != m0.Id ||
		m.Name != m0.Name ||
		m.Artist != m0.Artist ||
		m.Source != m0.Source ||
		m.Type != m0.Type {
		t.Error("m and m0 equal failed")
	}
	m, err := mm.Get(0)
	if err != nil {
		t.Error("get failed")
	}
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("remove failed")
	}
}
