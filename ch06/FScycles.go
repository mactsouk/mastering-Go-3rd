package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var visited = map[string]int{}

func walkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil
	}

	fileInfo, _ = os.Lstat(path)
	mode := fileInfo.Mode()

	// Find regular directories first
	if mode.IsDir() {
		abs, _ := filepath.Abs(path)
		_, ok := visited[abs]
		if ok {
			fmt.Println("Found cycle:", abs)
			return nil
		}
		visited[abs]++
		return nil
	}

	// Find symbolic links to directories
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		temp, err := os.Readlink(path)
		if err != nil {
			fmt.Println("os.Readlink():", err)
			return err
		}

		newPath, err := filepath.EvalSymlinks(temp)
		if err != nil {
			return nil
		}

		linkFileInfo, err := os.Stat(newPath)
		if err != nil {
			return err
		}

		linkMode := linkFileInfo.Mode()
		if linkMode.IsDir() {
			fmt.Println("Following...", path, "-->", newPath)
			abs, _ := filepath.Abs(newPath)
			_, ok := visited[abs]
			if ok {
				fmt.Println("Found cycle!", abs)
				return nil
			}
			visited[abs]++

			err = filepath.Walk(newPath, walkFunction)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}

	Path := arguments[1]
	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range visited {
		if v > 1 {
			fmt.Println(k, v)
		}
	}
}
