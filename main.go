package main

import (
	"github.com/sandro-lex-symphony/gobot"
	"github.com/sandro-lex-symphony/snykbot/processors"
)

func main() {
	gobot.Init()
	gobot.Loop(processors.Experiment)
}
