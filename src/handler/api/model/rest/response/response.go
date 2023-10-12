package response

type ResponseData[T any] struct {
	Error   []map[string]string `json:"error,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    T                   `json:"data,omitempty"`
}
