package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}

//implement Player interface
func (p *MP3Player) Play(source string) {
	fmt.Println("Mp3 Player")
	p.progress = 0
	for p.progress < 100 {
		//TODO realPlayer
		time.Sleep(100 * time.Millisecond) //
		fmt.Println(".....")
		p.progress += 10
	}
	fmt.Println(source, "Finish")
}
