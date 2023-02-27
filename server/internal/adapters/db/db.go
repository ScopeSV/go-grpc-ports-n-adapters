package db

import (
	"fmt"

	"github.com/scopesv/todoGrpc/server/internal/application/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func (a *Adapter) Save(todo *domain.Todo) (int64, error) {
	res := a.db.Save(&todo)
	if res.Error != nil {
		return -0, fmt.Errorf("Could not save todo: %v", res.Error)
	}

	return todo.ID, nil
}

func (a *Adapter) GetAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	res := a.db.Find(&todos)

	if res.Error != nil {
		return nil, fmt.Errorf("Could not get all todos: %v", res.Error)
	}

	return todos, nil
}

func (a *Adapter) Delete(id int64) error {
	res := a.db.Delete(&domain.Todo{}, id)

	fmt.Println(res.RowsAffected)

	if res.Error != nil {
		return fmt.Errorf("Could not delete todo: %v", res.Error)
	}

	return nil

}

func (a *Adapter) SetCompleted(id int64) error {
	res := a.db.Model(&domain.Todo{}).Where("id = ?", id).Update("completed", true)

	if res.Error != nil {
		return fmt.Errorf("Could not set todo as completed: %v", res.Error)
	}

	return nil
}

type TodoSchema struct {
	gorm.Model
	ID        int64  `json:"id"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
	CreatedAt int64  `json:"createdAt"`
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})

	if openErr != nil {
		return nil, fmt.Errorf("Could not connect to db: %v", openErr)
	}

	err := db.AutoMigrate(&TodoSchema{})

	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}
