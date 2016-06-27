package logging

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/soulski/dmp-cli"
)

func DMP(w http.ResponseWriter, r *http.Request, c *Context) {
	bodyBuff, err := ioutil.ReadAll(r.Body)

	if err != nil {
		c.RenderError(w, "Cannot read request", err)
		return
	}

	msg, err := dmpc.ParseMsg(bodyBuff)
	if err != nil {
		c.RenderError(w, "Msg content is invalid", err)
		return
	}

	switch msg.MsgType {
	case "create-log":
		var log *Log
		err := msg.ParseBody(&log)
		if err != nil {
			c.RenderError(w, "Log content is invalid", err)
			return
		}

		CreateLog(w, c, log)
	case "list-log":
		ListLog(w, c)

	default:
		c.RenderError(w, "Unknow type of msg", fmt.Errorf(""))
	}
}
