package engine

import (
	"fmt"
	"strings"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type printCommand struct {
	arg string
}

func (p *printCommand) Execute(loop Handler) {
	fmt.Println(p.arg)
}

type catCommand struct {
	arg1, arg2 string
}

func (mul *catCommand) Execute(loop Handler) {
	loop.Post(&printCommand{arg: mul.arg1 + mul.arg2})
}

func Parse(commandLine string) Command {
	parts := strings.Fields(commandLine)
	if len(parts) == 0 {
		return nil
	}
	if parts[0] == "print" && len(parts) > 1 {
		return &printCommand{arg: strings.Join(parts[1:], " ")}
	}
	if parts[0] == "cat" && len(parts) == 3 {
		return &catCommand{arg1: parts[1], arg2: parts[2]}
	}
	return &printCommand{arg: fmt.Sprintf("SYNTAX ERROR: command %s not valid", parts[0])}
}

type EventLoop struct {
	messageQueue []Command
}

func (e *EventLoop) Start() {
	e.messageQueue = make([]Command, 0)
}

func (e *EventLoop) Post(c Command) {
	if c != nil {
		e.messageQueue = append(e.messageQueue, c)
	}
}

func (e *EventLoop) AwaitFinish() {
	for i := 0; i < len(e.messageQueue); i++ {
		command := e.messageQueue[i]
		command.Execute(e)
	}
}
