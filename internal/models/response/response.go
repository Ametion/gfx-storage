package response

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
