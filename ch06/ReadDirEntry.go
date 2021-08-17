package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetSize(path string) (int64, error) {
	contents, err := os.ReadDir(path)
	if err != nil {
		return -1, err
	}

	var total int64
	for _, entry := range contents {
		// Visit directory entries
		if entry.IsDir() {
			temp, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return -1, err
			}
			total += temp
			// Get size of each non-directory entry
		} else {
			info, err := entry.Info()
			if err != nil {
				return -1, err
			}
			// Returns an int64 value
			total += info.Size()
		}
	}
	return total, nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a <Directory>")
		return
	}

	root, err := filepath.EvalSymlinks(arguments[1])
	fileInfo, err := os.Stat(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileInfo, _ = os.Lstat(root)
	mode := fileInfo.Mode()
	if !mode.IsDir() {
		fmt.Println(root, "not a directory!")
		return
	}

	i, err := GetSize(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total Size:", i)
}
