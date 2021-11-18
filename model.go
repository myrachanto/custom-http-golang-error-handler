package httperors


////////////errors ////////////////////////
type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

////////////success ////////////////////////
type HttpSuccess struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}
