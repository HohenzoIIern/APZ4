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

