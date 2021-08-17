package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type emptyFile struct {
	Ended bool
	Read  int
}

// Implement error interface
func (e emptyFile) Error() string {
	return fmt.Sprintf("Ended with io.EOF (%t) but read (%d) bytes", e.Ended, e.Read)
}

// Check values
func isFileEmpty(e error) bool {
	// Type assertion
	v, ok := e.(emptyFile)
	if ok {
		if v.Read == 0 && v.Ended == true {
			return true
		}
	}
	return false
}

func readFile(file string) error {
	var err error
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	n := 0
	for {
		line, err := reader.ReadString('\n')
		n += len(line)
		if err == io.EOF {
			// End of File: nothing more to read
			if n == 0 {
				return emptyFile{true, n}
			}
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("usage: errorInt <file1> [<file2> ...]")
		return
	}

	for _, file := range flag.Args() {
		err := readFile(file)
		if isFileEmpty(err) {
			fmt.Println(file, err)
		} else if err != nil {
			fmt.Println(file, err)
		} else {
			fmt.Println(file, "is OK.")
		}
	}
}
