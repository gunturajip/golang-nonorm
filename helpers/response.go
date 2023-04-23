package helpers

// Response struct includes the status code, success message (if exist),
// error message (if exist), and appropriate data (if exist) of a certain response
type Response struct {
	Status  int         `json:"status"`
	Success string      `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}
