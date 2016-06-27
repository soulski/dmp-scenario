package gateway

import (
	"io/ioutil"
	"net/http"

	"github.com/soulski/dmp-cli"
	"github.com/soulski/dmp-scenario/authen/server"
	"github.com/soulski/dmp-scenario/store/server"
	"github.com/soulski/dmp-scenario/user/server"
)

func CreateUser(w http.ResponseWriter, r *http.Request, c *Context) {
	bodyBuf, _ := ioutil.ReadAll(r.Body)

	msg, err := dmpc.NewMsgWithBuff("create-user", bodyBuf)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err := dmpc.Request("user", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	var user *user.User
	if err := c.ParseResult(res, &user); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	msg, err = dmpc.NewMsgWithBuff("create-authen", bodyBuf)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err = dmpc.Request("authen", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	var authen *authen.Authen
	if err := c.ParseResult(res, &authen); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	c.Render.JSON(w, http.StatusOK, user)
}

func CreateItem(w http.ResponseWriter, r *http.Request, c *Context) {
	bodyBuf, _ := ioutil.ReadAll(r.Body)

	msg, err := dmpc.NewMsgWithBuff("create-item", bodyBuf)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err := dmpc.Request("store", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	var item *store.Item
	if err := c.ParseResult(res, &item); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	c.Render.JSON(w, http.StatusOK, item)
}
