package todo

type Model struct {
	Id int
	Message string `json:"message,omitempty"`
}
