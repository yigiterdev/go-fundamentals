package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(400 * time.Millisecond)
	}
	channel <- "Done!"
}

func main() {
	var channel chan string
	go printMessage("Go is great!", channel)
	fmt.Println(<-channel)
	close(channel)
	// go printMessage("Yes it is...")
	// go printMessage("But JS...")
}
