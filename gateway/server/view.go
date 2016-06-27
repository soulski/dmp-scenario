package gateway

import (
	"fmt"
	"net/http"

	"github.com/soulski/dmp-cli"
	"github.com/soulski/dmp-scenario/analyse/server"
	"github.com/soulski/dmp-scenario/authen/server"
	"github.com/soulski/dmp-scenario/logging/server"
	"github.com/soulski/dmp-scenario/search/server"
	"github.com/soulski/dmp-scenario/store/server"
	"github.com/soulski/dmp-scenario/user/server"
)

type View struct {
	Title string
	Error error
	Data  interface{}
}

func Member(w http.ResponseWriter, r *http.Request, c *Context) {
	members, err := dmpc.GetAllMembers()
	if err != nil {
		c.Render.HTML(w, http.StatusOK, "error", &View{
			Title: "Error",
			Error: err,
		})
	} else {
		err := c.Render.HTML(w, http.StatusOK, "member", &View{
			Title: "Members",
			Data:  members,
		})
		if err != nil {
			fmt.Println("Erro : " + err.Error())
		}
	}
}

func User(w http.ResponseWriter, r *http.Request, c *Context) {
	//----------------------- list-user --------------------------------/
	msg, err := dmpc.NewMsg("list-user", nil)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err := dmpc.Request("user", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	users := make([]*user.User, 0)
	if err := c.ParseResult(res, &users); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	//----------------------- list-store --------------------------------/
	msg, err = dmpc.NewMsg("list-item", nil)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err = dmpc.Request("store", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	items := make([]*store.Item, 0)
	if err := c.ParseResult(res, &items); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	//----------------------- list-authen --------------------------------//
	msg, err = dmpc.NewMsg("list-authen", nil)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err = dmpc.Request("authen", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	authens := make([]*authen.Authen, 0)
	if err := c.ParseResult(res, &authens); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	//----------------------- render --------------------------------//
	c.Render.HTML(w, http.StatusOK, "user", &View{
		Title: "Users",
		Data: struct {
			Users   []*user.User
			Authens []*authen.Authen
			Items   []*store.Item
		}{users, authens, items},
	})
}

func Internal(w http.ResponseWriter, r *http.Request, c *Context) {
	//----------------------- list-event --------------------------------//
	msg, err := dmpc.NewMsg("list-event", nil)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err := dmpc.Request("analyse", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	events := make([]*analyse.Event, 0)
	if err := c.ParseResult(res, &events); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	//----------------------- list-search --------------------------------//
	msg, err = dmpc.NewMsg("list-record", nil)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err = dmpc.Request("search", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	records := make([]*search.Record, 0)
	if err := c.ParseResult(res, &records); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	//----------------------- render --------------------------------//
	c.Render.HTML(w, http.StatusOK, "internal", &View{
		Title: "Internal",
		Data: struct {
			Events  []*analyse.Event
			Records []*search.Record
		}{
			events,
			records,
		},
	})
}

func Log(w http.ResponseWriter, r *http.Request, c *Context) {
	msg, err := dmpc.NewMsg("list-log", nil)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	res, err := dmpc.Request("logging", msg)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	logs := make([]*logging.Log, 0)
	if err := c.ParseResult(res, &logs); err != nil {
		c.Render.JSON(w, http.StatusForbidden, err)
		return
	}

	c.Render.HTML(w, http.StatusOK, "log", &View{
		Title: "Log",
		Data: struct {
			Logs []*logging.Log
		}{
			logs,
		},
	})
}

func InternalError(w http.ResponseWriter, c *Context, err error) {
	c.Render.JSON(w, http.StatusForbidden, map[string]string{
		"ErrorCode": "3",
		"ErrorMsg":  "Operation fail",
	})

	fmt.Println("Error : ", err)
}
