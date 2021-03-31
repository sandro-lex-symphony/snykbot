package main

import (
	"fmt"

	"github.com/sandro-lex-symphony/gobot"
	"github.com/sandro-lex-symphony/snykbot/processors"
)

func main() {
	gobot.Init()
	fmt.Println("XXXXX")
	gobot.Loop(processors.Experiment)
}
