package main

import (
	"bufio"
	"fmt"
	library "goyard/simpleMediaPlayer/musicLibrary"
	"goyard/simpleMediaPlayer/musicPlayer"
	"os"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1

func handleLibCommand(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			music, _ := lib.Get(i)
			fmt.Println(i+1, ":", music.Name, music.Artist, music.Source, music.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{Id: strconv.Itoa(id), Name: tokens[2], Artist: tokens[3], Source: tokens[4], Type: tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name> <artist> <source> <type>")
		}
	case "remove": //TODO: 有Bug，删除不成功
		if len(tokens) == 3 {
			musicTemp := lib.Find(tokens[2])
			if musicTemp == nil {
				fmt.Println("The music ", tokens[2], "does not exist.")
				return
			}
			index, _ := strconv.Atoi(musicTemp.Id)
			_ = lib.Remove(index)
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command: ", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	music := lib.Find(tokens[1])
	if music == nil {
		fmt.Println("The music ", tokens[1], "does not exist.")
		return
	}
	musicPlayer.Play(music.Source, music.Type)
}

func main() {
	fmt.Println("Enter following commands to control the player: " +
		"lib list -- View the existing music lib" +
		"lib add <name><artist><source><type> -- Add a music to the music lib" +
		"lib remove <name> -- Remove the specified music from the lib" +
		"play <name> -- Play the specified music")
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter Command -> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommand(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
