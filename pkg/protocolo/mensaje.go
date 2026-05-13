package protocolo

// Mensaje es el "molde" JSON que viaja por los sockets
type Mensaje struct {
	Usuario string `json:"usuario"`
	Texto   string `json:"texto"`
}