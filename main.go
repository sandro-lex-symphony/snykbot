package main

import (
	"fmt"

	"github.com/sandro-lex-symphony/gobot"
	"github.com/sandro-lex-symphony/snykbot/processor"
)

func main() {
	gobot.Init()
	fmt.Println("XXXXX")
	gobot.Loop(processor.Experiment)
}
