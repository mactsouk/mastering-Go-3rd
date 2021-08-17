package main

import "fmt"

//go:generate ./echo.sh

//go:generate echo GOFILE: $GOFILE
//go:generate echo GOARCH: $GOARCH
//go:generate echo GOOS: $GOOS
//go:generate echo GOLINE: $GOLINE
//go:generate echo GOPACKAGE: $GOPACKAGE
//go:generate echo DOLLAR: $DOLLAR

//go:generate echo Hello!
//go:generate ls -l
//go:generate ./hello.py

func main() {
	fmt.Println("Hello there!")
}
