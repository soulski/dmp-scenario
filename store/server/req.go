package store

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type Log struct {
	Addr      string `json:"address"`
	Namespace string `json:"namespace"`
	Cause     string `json:"cause"`
}
