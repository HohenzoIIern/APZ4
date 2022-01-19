package main

import (
	"bufio"
	"os"

	"github.com/HohenzoIIern/APZ234/Lab4/engine"
)

// cat <arg1> <arg2>

const (
	inputFile = "commandList.txt"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine) // parse the line to get a Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
