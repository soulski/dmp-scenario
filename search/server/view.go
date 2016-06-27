package search

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
	case "create-user":
		var user *User
		if err := msg.ParseBody(&user); err != nil {
			c.RenderError(w, "user content is invalid", err)
			return
		}

		CreateRecord(w, c, &Record{
			IDRef: user.ID,
			Token: user.Username + " " + user.Email,
		})

	case "create-item":
		var item *Item
		if err := msg.ParseBody(&item); err != nil {
			c.RenderError(w, "item content is invalid", err)
			return
		}

		CreateRecord(w, c, &Record{
			IDRef: item.ID,
			Token: item.Name + " " + item.Description,
		})

	case "list-record":
		ListRecord(w, c)

	default:
		c.RenderError(w, "Unknow type of msg", fmt.Errorf(""))
	}
}

func CreateRecord(w http.ResponseWriter, c *Context, record *Record) {
	err := c.DB.InsertRecord(record.IDRef, record.Token)
	if err != nil {
		c.RenderError(w, "Operation fail", err)
		return
	}

	c.RenderJSON(w, record)
}

func ListRecord(w http.ResponseWriter, c *Context) {
	records, err := c.DB.ListAllRecord()
	if err != nil {
		c.RenderError(w, "Operation fail", err)
		return
	}
	c.RenderJSON(w, records)
}
