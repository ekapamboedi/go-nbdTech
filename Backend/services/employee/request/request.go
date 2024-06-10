package request

type RequestCreateEmployee struct {
	Id      int64  `jsom:"Id"`
	Name    string `jsom:"Name"`
	Email   string `jsom:"Email"`
	Address string `jsom:"Address"`
	Phone   string `jsom:"Phone"`
}

type RequestUpdate struct {
	Name    string `jsom:"Name"`
	Email   string `jsom:"Email"`
	Address string `jsom:"Address"`
	Phone   string `jsom:"Phone"`
}
