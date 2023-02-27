package domain

type Todo struct {
	ID        int64  `json:"id"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
	CreatedAt int64  `json:"createdAt"`
}

type TodoRequest struct {
	Desc string `json:"desc"`
}
