package main

import (
	"fmt"
)

func main() {
	willClose := make(chan complex64, 10)

	// Write some data to the channel
	willClose <- -1
	willClose <- 1i

	// Read data and empty channel
	<-willClose
	<-willClose
	close(willClose)

	// Read again – this is a closed channel
	read := <-willClose
	fmt.Println(read)
}
