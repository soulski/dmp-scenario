package authen

import (
	"encoding/json"
)

type Msg struct {
	MsgType string          `json:"type"`
	Body    json.RawMessage `json:"body"`
}

func (m *Msg) ParseBody(bType interface{}) (err error) {
	return json.Unmarshal(m.Body, bType)
}

type Authen struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
