package analyse

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
	case "create-event":
		var event *Event
		if err := msg.ParseBody(&event); err != nil {
			c.RenderError(w, "Authen content is invalid", err)
			return
		}

		CreateEvent(w, c, event)

	case "list-event":
		ListEvent(w, c)

	case "create-user":
		var user *User
		if err := msg.ParseBody(&user); err != nil {
			c.RenderError(w, "User content is invalid", err)
			return
		}

		CreateEvent(w, c, &Event{
			TargetID: user.ID,
			Event:    "create-user",
			Time:     time.Now().UnixNano(),
		})

	case "create-item":
		var item *Item
		if err := msg.ParseBody(&item); err != nil {
			c.RenderError(w, "Item content is invalid", err)
			return
		}

		CreateEvent(w, c, &Event{
			TargetID: item.ID,
			Event:    "create-item",
			Time:     time.Now().UnixNano(),
		})

	default:
		c.RenderError(w, "Unknow type of msg", fmt.Errorf(""))
	}
}

func CreateEvent(w http.ResponseWriter, c *Context, event *Event) {
	err := c.DB.InsertEvent(event.TargetID, event.Event, event.Time)
	if err != nil {
		c.RenderError(w, "Operation fail", err)
		return
	}

	c.RenderJSON(w, event)
}

func ListEvent(w http.ResponseWriter, c *Context) {
	events, err := c.DB.ListAllEvent()
	if err != nil {
		c.RenderError(w, "Operation fail", err)
		return
	}

	c.RenderJSON(w, events)
}
