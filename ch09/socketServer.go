package main

import (
	"fmt"
	"net"
	"os"
)

func echo(c net.Conn) {
	for {
		buf := make([]byte, 128)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Read:", err)
			return
		}

		data := buf[0:n]
		fmt.Print("Server got: ", string(data))
		_, err = c.Write(data)
		if err != nil {
			fmt.Println("Write:", err)
			return
		}
	}
}

func main() {
	// Read socket path
	if len(os.Args) == 1 {
		fmt.Println("Need socket path")
		return
	}
	socketPath := os.Args[1]

	// If socketPath exists, delete it
	_, err := os.Stat(socketPath)
	if err == nil {
		fmt.Println("Deleting existing", socketPath)
		err := os.Remove(socketPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			return
		}
		go echo(fd)
	}
}
