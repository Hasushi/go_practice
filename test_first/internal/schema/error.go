package schema

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Errors 	    []string `json:"errors,omitempty"`
}