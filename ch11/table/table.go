package table

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var BUFFERSIZE int
var FILESIZE int

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func CreateBuffer(buf *[]byte, count int) {
	*buf = make([]byte, count)
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		intByte := byte(random(50, 100))
		if len(*buf) > count {
			return
		}
		*buf = append(*buf, intByte)
	}
}

func Create(dst string, b, filesize int) error {
	_, err := os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	f, err := os.Create(dst)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		return err
	}
	f.Close()

	f, err = os.OpenFile(dst, os.O_WRONLY, 0655)
	if err != nil {
		return err
	}
	defer f.Close()

	bwriter := bufio.NewWriterSize(f, b)
	if err != nil {
		return err
	}

	buf := make([]byte, 0)
	CreateBuffer(&buf, b)
	buf = buf[:b]
	for {
		_, err := bwriter.Write(buf)
		if err != nil {
			return err
		}

		if filesize < 0 {
			break
		}
		filesize = filesize - len(buf)
	}
	return err
}

func CountChars(filename string, b int) int {
	buf := make([]byte, 0)
	CreateBuffer(&buf, b)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		return -1
	}
	defer f.Close()

	size := 0
	for {
		n, err := f.Read(buf)
		size = size + n
		if err != nil {
			break
		}
	}

	return size
}
