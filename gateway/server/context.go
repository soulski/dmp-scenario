package gateway

import (
	"fmt"
	"net/http"

	"github.com/soulski/dmp-cli"
	"github.com/unrolled/render"
)

type Context struct {
	Render *render.Render
}

func NewContext(templatePath string) *Context {
	r := render.New(render.Options{
		Directory:     templatePath,
		Layout:        "layout",
		IsDevelopment: true,
	})

	return &Context{
		Render: r,
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

func (c *Context) ParseResult(buffResult []byte, msg interface{}) *dmpc.Error {
	result, err := dmpc.ParseResult(buffResult)
	if err != nil {
		fmt.Println("Error : ")
		fmt.Println(string(buffResult))
		panic(err)
	}

	if !result.Action {
		return result.Error
	}

	err = result.ParseMsg(msg)
	if err != nil {
		fmt.Println("Error : ")
		fmt.Println(string(result.Message))
		panic(err)
	}

	return nil
}
