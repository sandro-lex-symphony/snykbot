package processors

import (
	"github.com/sandro-lex-symphony/gobot"
	snykTool "github.com/sandro-lex-symphony/snykTools"
)

func Experiment(sid, data string) {
	//cmd := gobot.GetCommand(data)
	//gobot.Fatal(err)

	if gobot.IsMentioned(data) {
		result, err := snykTool.GetOrgs()
		gobot.Fatal(err)
		var msg string
		msg = "<card  accent='tempo-bg-color--blue' iconSrc=''><body>"
		for _, org := range result.Orgs {
			msg += "<p><b>" + org.Name + "</b> </p>"
		}
		msg += "</body></card>"
		err = gobot.SendMessage(sid, msg)
		gobot.Fatal(err)
	}

}
