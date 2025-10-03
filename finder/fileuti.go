package finder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// buildPath full path  to the file to be created. If the filePath has no parent directory 
// buildPath returns it, otherwise, it returns the full path created with renof-docs, ready 
// for IO operations.
func buildPath(filePath string) (string, error) {
	var parents string
	if strings.Contains(filePath, string(os.PathSeparator)) {
		parents = filepath.Dir(filePath)
	} else {
		return filePath, nil
	}
	if len(parents) != 0 && strings.Contains(parents, HOME_DIR) {
		// to remove leading '/', otherwise os.error.20
		_, p, _ := strings.Cut(parents, HOME_DIR + string(os.PathSeparator))
		parents = p
		fmt.Println("[*] home_dir stripped ->", parents)
	}

	lastOccurence := strings.LastIndex(filePath, string(os.PathSeparator))
	if lastOccurence == -1 {
		log.Print()
		return "", fmt.Errorf("[x] could not find last seperator")
	}
	src := filePath[lastOccurence:]
	fmt.Printf("[*] creating parents -> %s\n", filepath.Join(BASE_DIR, parents, src))
	if err := os.MkdirAll(filepath.Join(BASE_DIR, parents), 0777); err != nil {
		return "", fmt.Errorf("[x] error occured creating parent directories -> %w", err)
	}
	return filepath.Join(BASE_DIR, parents, src), nil
}
