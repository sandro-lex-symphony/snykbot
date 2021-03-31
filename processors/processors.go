package processors

import (
	"fmt"

	"github.com/sandro-lex-symphony/gobot"
	snykTool "github.com/sandro-lex-symphony/snykTools"
)

func Hello(sid, data string) {
	err := gobot.SendMessage(sid, "<card accent='tempo-bg-color--blue' iconSrc=''><body><h1>filler</h1></body></card>")
	gobot.Fatal(err)
}

func Experiment(sid, data string) {
	cmd := gobot.GetCommand(data)
	err := gobot.SendMessage(sid, cmd)
	gobot.Fatal(err)

	fmt.Println("KJFLKJLD")
	result, err := snykTool.GetOrgs()
	gobot.Fatal(err)
	var msg string
	msg = "<card  accent='tempo-bg-color--blue' iconSrc=''><body>"
	for _, org := range result.Orgs {
		msg += "<p><b>" + org.Name + "</b> </p>"
	}
	msg += "</body></card>"
	fmt.Println(msg)
	err = gobot.SendMessage(sid, msg)
	gobot.Fatal(err)
}
