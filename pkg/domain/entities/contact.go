package entities

type Contact struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Deleted bool   `json:"deleted"`
}
