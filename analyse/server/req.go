package analyse

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

type Msg struct {
	MsgType string          `json:"type"`
	Body    json.RawMessage `json:"body"`
}

func NewMsgWithReader(mType string, r io.Reader) (*Msg, error) {
	bodyBuff, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return &Msg{
		MsgType: mType,
		Body:    bodyBuff,
	}, nil
}

func (m *Msg) ParseBody(bType interface{}) (err error) {
	return json.Unmarshal(m.Body, bType)
}

type Event struct {
	TargetID int    `json:"targetID"`
	Event    string `json:"event"`
	Time     int64  `json:"time"`
}

func (e *Event) DisplayTime() string {
	return time.Unix(e.Time, 0).String()
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
