package main

import (
	"bufio"
	"fmt"
	"io"
)

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1   S1
	text []byte
}

// Using pointer to S1 for changes to be persistent when the method exits
func (s *S1) Read(p []byte) (n int, err error) {
	fmt.Print("Give me your name: ")
	fmt.Scanln(&p)
	s.F2 = string(p)
	return len(p), nil
}

func (s *S1) Write(p []byte) (n int, err error) {
	if s.F1 < 0 {
		return -1, nil
	}

	for i := 0; i < s.F1; i++ {
		fmt.Printf("%s ", p)
	}
	fmt.Println()
	return s.F1, nil
}

func (s S2) eof() bool {
	return len(s.text) == 0
}

func (s *S2) readByte() byte {
	// this function assumes that eof() check was done before
	temp := s.text[0]
	s.text = s.text[1:]
	return temp
}

func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		err = io.EOF
		return
	}

	l := len(p)
	if l > 0 {
		for n < l {
			p[n] = s.readByte()
			n++
			if s.eof() {
				s.text = s.text[0:0]
				break
			}
		}
	}
	return
}

func main() {
	s1var := S1{4, "Hello"}
	fmt.Println(s1var)

	buf := make([]byte, 2)
	_, err := s1var.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Read:", s1var.F2)
	_, _ = s1var.Write([]byte("Hello There!"))

	s2var := S2{F1: s1var, text: []byte("Hello world!!")}
	// Read s2var.text
	r := bufio.NewReader(&s2var)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("*", err)
			break
		}
		fmt.Println("**", n, string(buf[:n]))
	}
}
