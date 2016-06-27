package store

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
	case "create-item":
		var item *Item
		if err := msg.ParseBody(&item); err != nil {
			LogError(c, err)
			c.RenderError(w, "Item content is invalid", err)
			return
		}

		CreateItem(w, c, item)
	case "list-item":
		ListItem(w, c)

	default:
		c.RenderError(w, "Unknow type of msg", fmt.Errorf(""))
	}
}

func CreateItem(w http.ResponseWriter, c *Context, item *Item) {
	item, err := c.DB.InsertItem(item.Name, item.Description, item.Price)

	if err != nil {
		LogError(c, err)
		InternalError(w, c, err)
		return
	}

	body, err := dmpc.NewMsg("create-item", item)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	_, err = dmpc.Publish("create-item", body)
	if err != nil {
		LogError(c, err)
		InternalError(w, c, err)
		return
	}

	c.RenderJSON(w, item)
}

func ListItem(w http.ResponseWriter, c *Context) {
	items, err := c.DB.ListAllItem()
	if err != nil {
		InternalError(w, c, err)
	} else {
		c.RenderJSON(w, items)
	}
}

func InternalError(w http.ResponseWriter, c *Context, err error) {
	c.RenderError(w, "Operation fail", err)
	fmt.Println("Error : ", err)
}

func LogError(c *Context, err error) {
	log := &Log{
		Addr:      c.IP.String(),
		Namespace: c.Namespace,
		Cause:     err.Error(),
	}

	msg, err := dmpc.NewMsg("create-log", log)
	if err != nil {
		fmt.Println("Error : Unable to send error to logging system")
		fmt.Println(err)
		return
	}

	_, err = dmpc.Notificate("logging", msg)
	if err != nil {
		fmt.Println("Error : Unable to send error to logging system")
		fmt.Println(err)
	}
}
