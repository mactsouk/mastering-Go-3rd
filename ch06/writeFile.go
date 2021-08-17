package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	buffer := []byte("Data to write\n")

	f1, err := os.Create("/tmp/f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(buffer))

	f2, err := os.Create("/tmp/f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(buffer))
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("/tmp/f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(buffer))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	f := "/tmp/f4.txt"
	f4, err := os.Create(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()

	for i := 0; i < 5; i++ {
		n, err = io.WriteString(f4, string(buffer))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("wrote %d bytes\n", n)
	}

	// Append to a file
	f4, err = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()

	// Write() needs a byte slice
	n, err = f4.Write([]byte("Put some more data at the end.\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
