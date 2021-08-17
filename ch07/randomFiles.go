package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createFile(file string) {
	// Does the file already exist?
	_, err := os.Stat(file)
	// File already exist - do not override
	if err == nil {
		fmt.Printf("%s already exists!\n", file)
		return
	}

	// Can the file be created?
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// The number of lines is randomly chosen
	lines := random(10, 30)
	// Populate file with random data
	for i := 0; i < lines; i++ {
		data := random(0, 20)
		fmt.Fprintf(f, "%d\n", data)
	}
	fmt.Printf("%s created!\n", file)
}

func main() {
	arguments := os.Args
	if len(arguments) != 5 {
		fmt.Println("Usage: randomFiles firstInt lastInt filename directory")
		return
	}

	start, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	end, err := strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	// For the for loop to work properly
	if end < start {
		fmt.Println(end, "<", start)
		return
	}

	filename := arguments[3]
	path := arguments[4]
	// Does the destination directory exist?
	_, err = os.Open(path)
	if err != nil {
		fmt.Println(path, "does not exist!")
		return
	}

	var waitGroup sync.WaitGroup
	for i := start; i <= end; i++ {
		waitGroup.Add(1)
		filepath := fmt.Sprintf("%s/%s%d", path, filename, i)
		go func(f string) {
			defer waitGroup.Done()
			createFile(f)
		}(filepath)
	}
	waitGroup.Wait()
}
