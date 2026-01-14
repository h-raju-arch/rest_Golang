package types

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TodoOutput struct {
	Id          int
	Title       string
	Description string
	Done        bool
}
