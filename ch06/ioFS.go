package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed static
var f embed.FS

var searchString string

func walkFunction(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("Path=%q, isDir=%v\n", path, d.IsDir())
	return nil
}

func walkSearch(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.Name() == searchString {
		fileInfo, err := fs.Stat(f, path)
		if err != nil {
			return err
		}
		fmt.Println("Found", path, "with size", fileInfo.Size())
		return nil
	}
	return nil
}

func list(f embed.FS) error {
	return fs.WalkDir(f, ".", walkFunction)
}

func search(f embed.FS) error {
	return fs.WalkDir(f, ".", walkSearch)
}

func extract(f embed.FS, filepath string) ([]byte, error) {
	s, err := fs.ReadFile(f, filepath)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func writeToFile(s []byte, path string) error {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	n, err := fd.Write(s)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n)
	return nil
}

func main() {
	// At this point we do not know what is included in ./static

	// List all files
	err := list(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Search
	searchString = "file.txt"
	err = search(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract into a byte slice
	buffer, err := extract(f, "static/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save it to an actual file
	writeToFile(buffer, "/tmp/IOFS.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
