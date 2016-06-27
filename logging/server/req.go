package logging

type Log struct {
	ID        int    `json:"id"`
	Addr      string `json:"address"`
	Namespace string `json:"namespace"`
	Cause     string `json:"cause"`
}
