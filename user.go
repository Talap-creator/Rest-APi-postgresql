package todo

type Todolist struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Userlist struct {
	Id     int
	UserID int
	ListID int
}

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	UserID int
	ListID int
}
