package finder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/persona-mp3/renof/soundex"
)

type Soundenc struct {
	Name string // this will be removed later on, used for debugging 
	Enc  string
}
type Location struct {
	string
}

var store = make(map[Soundenc]Location)
var man = "man"
var contrib = "contribute"

const (
	BASE_DIR = "renof-docs"
)

func LoadDefaults() {
	soundec, contribSoundec := Soundenc{
		Name: man,
		Enc:  soundex.Soundex(man),
	}, Soundenc{
		Name: contrib,
		Enc:  soundex.Soundex(contrib),
	}
	manLocation, contribLocation := Location{filepath.Join(BASE_DIR, man)}, Location{filepath.Join(BASE_DIR, contrib)}

	store[soundec] = manLocation
	store[contribSoundec] = contribLocation

	fmt.Println("[*] default configs loaded")
}

func (s *Soundenc) Add(location string) bool {
	_, exists := store[*s]

	var u string
	if !exists {
		fmt.Printf("[i] entry for %s already exists, update?[y/n]: ", s.Name)
		_, err := fmt.Scan(&u)
		if err != nil {
			panic(err)
		}

		if strings.ToLower(u) == "n" {
			return false
		} else {
			fmt.Printf("[*] existing entry for %s updated\n", s.Name)
		}
	}
	store[*s] = Location{filepath.Join(BASE_DIR, location)}
	// TODO: Create file and its actual content
	return true
}

func Get(s string) (string, error) {
	var enc Soundenc
	soundenc := soundex.Soundex(s)
	enc.Name = s
	enc.Enc = soundenc

	location, exists := store[enc]
	if !exists {
		return "", fmt.Errorf("[x] could not find entry in store")
	}

	return location.string, nil
}

func Display(path string) error {
	// Later on, we'll change this to match the users default CAT or LESS
	// As we'd have to search through user's path if either are configured
	cmd := exec.Command("bat", path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("[x] error occured while running command")
		return err
	}

	return nil
}

