package finder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Returns full path  to the file to be created
func buildPath(filePath string) (string, error) {
	var parents string
	if strings.Contains(filePath, string(os.PathSeparator)) {
		parents = filepath.Dir(filePath)
	} else {
		return filePath, nil
	}
	fmt.Println("[*] extracted parents -> ", parents)
	// remove home-dir
	if len(parents) != 0 && strings.Contains(parents, HOME_DIR) {
		_, p, _ := strings.Cut(parents, HOME_DIR+string(os.PathSeparator))
		parents = p
		fmt.Println("[*] home_dir stripped ->", parents)
	}

	lastOccurence := strings.LastIndex(filePath, string(os.PathSeparator))
	if lastOccurence == -1 {
		log.Print()
		return "", fmt.Errorf("[x] could not find last seperator")
	}
	src := filePath[lastOccurence:]
	fmt.Printf("[*] parent-directories -> %s, src-file -> %s, full-path -> %s\n", parents, src, filepath.Join(BASE_DIR, parents, src))
	if err := os.MkdirAll(filepath.Join(BASE_DIR, parents), 0777); err != nil {
		return "", fmt.Errorf("[x] error occured creating parent directories -> %w", err)
	}
	return filepath.Join(BASE_DIR, parents, src), nil
}
