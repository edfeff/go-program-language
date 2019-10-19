package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
}

//implement Player interface
func (p *WAVPlayer) Play(source string) {
	fmt.Println("wav Player")
	p.progress = 0
	for p.progress < 100 {
		//TODO realPlayer
		time.Sleep(100 * time.Millisecond) //
		fmt.Println(".....")
		p.progress += 10
	}
	fmt.Println(source, "Finish")
}
