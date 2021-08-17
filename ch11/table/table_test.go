package table

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"testing"
)

var ERR error
var countChars int

func benchmarkCreate(b *testing.B, buffer, filesize int) {
	filename := path.Join(os.TempDir(), strconv.Itoa(buffer))
	filename = filename + "-" + strconv.Itoa(filesize)
	var err error
	for i := 0; i < b.N; i++ {
		err = Create(filename, buffer, filesize)
	}
	ERR = err

	err = os.Remove(filename)
	if err != nil {
		fmt.Println(err)
	}
	ERR = err
}

func BenchmarkBuffer4Create(b *testing.B) {
	benchmarkCreate(b, 4, 1000000)
}

func BenchmarkBuffer8Create(b *testing.B) {
	benchmarkCreate(b, 8, 1000000)
}

func BenchmarkBuffer16Create(b *testing.B) {
	benchmarkCreate(b, 16, 1000000)
}

func BenchmarkRead(b *testing.B) {
	buffers := []int{1, 16, 96}
	files := []string{"10.txt", "1000.txt", "5k.txt"}

	for _, filename := range files {
		for _, bufSize := range buffers {
			name := fmt.Sprintf("%s-%d", filename, bufSize)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					t := CountChars("./testdata/"+filename, bufSize)
					countChars = t
				}
			})
		}
	}
}
