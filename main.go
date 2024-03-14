package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	instanceIdx := os.Getenv("CF_INSTANCE_INDEX")
	instanceID := os.Getenv("CF_INSTANCE_GUID")

	if len(instanceIdx) == 0 {
		instanceIdx = "0"
	}
	if len(instanceID) == 0 {
		instanceID = "(unknown)"
	}

	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	fmt.Printf("Waiting for exit signal - %s - %s\n", instanceIdx, instanceID)

	sig := <-sigCh
	switch sig {
	case os.Interrupt:
		fmt.Printf("Received SIGINT - %s - %s\n", instanceIdx, instanceID)
	case syscall.SIGTERM:
		fmt.Printf("Received SIGTERM - %s - %s\n", instanceIdx, instanceID)
		time.Sleep(9900)
		fmt.Printf("Done sleeping - %s - %s\n", instanceIdx, instanceID)
	case syscall.SIGKILL:
		fmt.Printf("Received SIGKILL - %s - %s\n", instanceIdx, instanceID)
	}
}
