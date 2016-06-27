package authen

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
	case "create-authen":
		var authen *Authen
		if err := msg.ParseBody(&authen); err != nil {
			c.RenderError(w, "Authen content is invalid", err)
			return
		}

		CreateAuthen(w, c, authen)
	case "login":
		var authen *Authen
		if err := msg.ParseBody(&authen); err != nil {
			c.RenderError(w, "Authen content is invalid", err)
			return
		}

		Login(w, c, authen)

	case "list-authen":
		ListAuthen(w, c)

	default:
		c.RenderError(w, "Unknow type of msg", fmt.Errorf(""))
	}
}

func CreateAuthen(w http.ResponseWriter, c *Context, authen *Authen) {
	err := c.DB.InsertAuthen(authen.Username, authen.Password)
	if err != nil {
		c.RenderError(w, "Operation fail", err)
		return
	}

	c.RenderJSON(w, authen)
}

func Login(w http.ResponseWriter, c *Context, authen *Authen) {
	authen, err := c.DB.QueryAuthen(authen.Username)
	if err != nil {
		c.RenderJSON(w, []byte("{ 'status': 'fail' }"))
		return
	}

	c.RenderJSON(w, []byte("{ 'status': 'success' }"))
}

func ListAuthen(w http.ResponseWriter, c *Context) {
	authen, err := c.DB.ListAllAuthen()
	if err != nil {
		c.RenderError(w, "Operation fail", err)
		return
	}

	c.RenderJSON(w, authen)
}
