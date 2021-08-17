package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) == 1 {
		return errors.New("No input!")
	}

	// Continue with the implementation of run()
	// as you would have with main()

	return nil
}
