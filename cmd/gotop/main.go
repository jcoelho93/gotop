package main

import (
	"fmt"
	"gotop/internal/pkg/display"
	"time"
)

func main() {
	fmt.Println("Starting gotop...")

	topDisplay := display.Top{}

	for {
		err := topDisplay.Display()
		if err != nil {
			fmt.Println("Error displaying stats:", err)
			return
		}
		time.Sleep(time.Second)
	}
}
