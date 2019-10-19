package mp

import "testing"

func TestMP3Player_Play(t *testing.T) {
	p := &MP3Player{}
	p.Play("hello")
}
