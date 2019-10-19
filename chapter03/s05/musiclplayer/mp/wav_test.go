package mp

import "testing"

func TestWAVPlayer_Play(t *testing.T) {
	p := &WAVPlayer{}
	p.Play("hello")
}
