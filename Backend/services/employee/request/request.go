package request

type RequestCreateEmployee struct {
	// Id      int64  `jsom:"Id"`
	Name    string `jsom:"Name"`
	Phone   string `jsom:"Phone"`
	Email   string `jsom:"Email"`
	Address string `jsom:"Address"`
}

type RequestUpdateEmployee struct {
	Name    string `jsom:"Name"`
	Phone   string `jsom:"Phone"`
	Email   string `jsom:"Email"`
	Address string `jsom:"Address"`
}
