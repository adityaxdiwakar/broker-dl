package main

import (
	"fmt"
	"log"

	"github.com/nxadm/tail"
)

func main() {
	t, err := tail.TailFile("log.txt", tail.Config{Follow: true})
	if err != nil {
		log.Fatalf("error: could not open file to tail from: %v\n", err)
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
