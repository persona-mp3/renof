package main

import (
	"fmt"
	"os"

	"github.com/persona-mp3/renof/finder"
	"github.com/persona-mp3/renof/soundex"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Enter value to encode")
		return
	}

	finder.LoadDefaults()
	if len(os.Args) > 3 && os.Args[1] == "add" {
		enc := soundex.Soundex(os.Args[1])
		s := &finder.Soundenc{Name: os.Args[2], Enc: enc}
		location := os.Args[3]
		ok, err := s.Add(location)
		if err != nil || !ok {
			fmt.Println(err)
			return
		}
		return
	}

	location, err := finder.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	finder.Display(location)
}
