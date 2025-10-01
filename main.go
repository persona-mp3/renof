package main

import (
	"fmt"
	"os"

	soundex "github.com/persona-mp3/renof/soundex"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Enter value to encode")
		return
	}
	soundex.SoundEnc(os.Args[1])
}
