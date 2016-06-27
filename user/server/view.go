package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

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
			c.RenderError(w, "User content is invalid", err)
			return
		}

		CreateUser(w, c, user)
	case "list-user":
		ListUser(w, c)

	default:
		c.RenderError(w, "Unknow type of msg", fmt.Errorf(""))
	}
}

func CreateUser(w http.ResponseWriter, c *Context, user *User) {
	re := regexp.MustCompile(".+@.+\\..+")
	validedEmail := re.Match([]byte(user.Email))

	if !validedEmail {
		err := fmt.Errorf("Email %s is invalid", user.Email)

		LogError(c, err)
		c.RenderError(w, "Bad arguments", err)
		return
	}

	user, err := c.DB.InsertUser(user.Username, user.Email, user.Address)

	if err != nil {
		LogError(c, err)
		InternalError(w, c, err)
		return
	}

	body, err := dmpc.NewMsg("create-user", user)
	if err != nil {
		InternalError(w, c, err)
		return
	}

	_, err = dmpc.Publish("create-user", body)
	if err != nil {
		LogError(c, err)
		InternalError(w, c, err)
		return
	}

	c.RenderJSON(w, user)
}

func ListUser(w http.ResponseWriter, c *Context) {
	users, err := c.DB.ListAllUser()
	if err != nil {
		InternalError(w, c, err)
	} else {
		c.RenderJSON(w, users)
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
