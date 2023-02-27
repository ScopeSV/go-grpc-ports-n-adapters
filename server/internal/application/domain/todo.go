package domain

import "time"

type Todo struct {
	ID        int64  `json:"id"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
	CreatedAt int64  `json:"createdAt"`
}

func NewTodo(desc string) *Todo {
	return &Todo{
		Desc:      desc,
		Completed: false,
		CreatedAt: time.Now().Unix(),
	}

}
