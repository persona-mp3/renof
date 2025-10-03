package finder

import (
	"fmt"
	"io"
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
	BASE_DIR  = "renof-docs"
	FILE_PERM = 0o744
)

var HOME_DIR string

func LoadDefaults() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	_, HOME_DIR, _ = strings.Cut(homeDir, string(os.PathSeparator))
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
	fmt.Println("[*] user $HOME_DIR configured -> ", HOME_DIR)
}

// Add() returns false if no IO operation occured. Panics if an error occured
// in collecting user-input. And returns an error based on IO operation
func (s *Soundenc) Add(location string) (bool, error) {
	_, exists := store[*s]

	var u string
	if exists {
		fmt.Printf("[i] entry for %s already exists, update?[y/n]: ", s.Name)
		_, err := fmt.Scan(&u)
		if err != nil {
			panic(err)
		}

		if strings.ToLower(u) == "n" {
			return false, nil
		} else {
			fmt.Printf("[*] existing entry for %s updated\n", s.Name)
		}
	}
	l := Location{filepath.Join(BASE_DIR, location)}
	if err := l.CreateFile(); err != nil {
		return false, err
	}
	store[*s] = l
	return true, nil
}

func (l Location) CreateFile() error {
	var dst string
	if strings.Contains(l.string, string(os.PathSeparator)) {
		p, err := buildPath(l.string)
		if err != nil {
			return err
		}
		dst = p
		fmt.Println("[**] built path -> ", dst)
	} else {
		fmt.Println("[*] file can be added to root")
		dst = l.string
	}

	// REEFACTOR!: Don't append the renof-docs, till all path is built
	_, src, ok := strings.Cut(l.string, string(os.PathSeparator))
	if !ok {
		return fmt.Errorf("[?] unexpected error occured getting source")
	}

	srcFile, err := os.Open(string(os.PathSeparator) + src)
	if err != nil {
		return fmt.Errorf("[x] error occured opening source file %w", err)
	}
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("[x] error occured while creating dst file -> %w", err)
	}

	io.Copy(dstFile, srcFile)
	fmt.Println("[✅] successfully created file")
	return nil
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
