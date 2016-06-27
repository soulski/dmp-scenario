package logging

import (
	"fmt"
	"net/http"
)

func CreateLog(w http.ResponseWriter, c *Context, log *Log) {
	log, err := c.DB.InsertLog(log.Addr, log.Namespace, log.Cause)
	if err != nil {
		fmt.Println("Error : insert log error")
		fmt.Println(err)
	}
}

func ListLog(w http.ResponseWriter, c *Context) {
	logs, err := c.DB.ListAllLog()
	if err != nil {
		c.RenderError(w, "Fail to read logs", err)
		return
	}

	c.RenderJSON(w, logs)
}
