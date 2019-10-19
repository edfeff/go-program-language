package main

import (
	"bufio"
	"fmt"
	"github.com/wpp/go-program-language/chapter03/s05/musiclplayer/library"
	"github.com/wpp/go-program-language/chapter03/s05/musiclplayer/mp"
	"os"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		listAllMusic()
	case "add":
		addMusic(tokens)
	case "remove":
		removeMusic(tokens)
	default:
		fmt.Println("operation is not allow")
	}
}
func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	music := lib.Find(tokens[1])
	if music == nil {
		fmt.Println("no this music")
		return
	}
	mp.Play(music.Source, music.Type)
	//TODO
	//mp.Play(music.Source, music.Type, ctrl, signal)
}

func removeMusic(tokens []string) {
	if len(tokens) == 3 {
		lib.RemoveByName(tokens[2])
	} else {
		fmt.Println("USAGE: lib add <name><artist><source><type>")
	}
}

func addMusic(tokens []string) {
	if len(tokens) == 6 {
		id++
		lib.Add(&library.MusicEntry{
			Id:     strconv.Itoa(id),
			Name:   tokens[2],
			Artist: tokens[3],
			Source: tokens[4],
			Type:   tokens[5],
		})
	} else {
		fmt.Println("USAGE: lib add <name><artist><source><type>")
	}
}

func listAllMusic() {
	for i := 0; i < lib.Len(); i++ {
		music, err := lib.Get(i)
		if err != nil {
			fmt.Println("music", i, "is error:", err)
		} else {
			fmt.Println("music：", music.Name, music.Type, music.Source, music.Artist)
		}
	}
}
func main() {
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music`)

	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
