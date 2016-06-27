package authen

import (
	"net/http"

	"github.com/soulski/dmp-cli"
	"github.com/unrolled/render"
)

type Context struct {
	Render *render.Render
	DB     *DB
}

func NewContext(db *DB) *Context {
	r := render.New()

	return &Context{
		Render: r,
		DB:     db,
	}
}

func (c *Context) RenderError(w http.ResponseWriter, msg string, err error) {
	result, err := dmpc.CreateErrorResult(&dmpc.Error{
		Message: msg,
		Cause:   err.Error(),
	})

	if err != nil {
		panic(err)
	}

	c.Render.Data(w, http.StatusOK, result)
}

func (c *Context) RenderJSON(w http.ResponseWriter, msg interface{}) {
	result, err := dmpc.CreateMsgResult(msg)

	if err != nil {
		panic(err)
	}

	c.Render.Data(w, http.StatusOK, result)
}
