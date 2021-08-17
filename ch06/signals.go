package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("handleSignal() Caught:", sig)
}

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	start := time.Now()
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Execution time:", duration)
			case syscall.SIGINFO:
				handleSignal(sig)
				// do not use return here because the goroutine will exit
				// but the time.Sleep() will continue to work!
				os.Exit(0)
			default:
				fmt.Println("Caught:", sig)
			}
		}
	}()

	for {
		fmt.Print("+")
		time.Sleep(10 * time.Second)
	}
}
