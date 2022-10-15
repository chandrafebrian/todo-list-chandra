package db

import "github.com/chandrafebrian/todo-list-chandra/tree/todo-master/backend/schema"

type Static struct{}

func (s *Static) GetAll() ([]schema.Todo, error) {
	todoList := []schema.Todo{
		{
			ID:   1,
			Note: "a",
			Done: false,
		},
		{
			ID:   2,
			Note: "b",
			Done: true,
		},
		{
			ID:   3,
			Note: "c",
			Done: false,
		},
	}
	return todoList, nil
}

func (s *Static) Insert(todo *schema.Todo) (int, error) {
	return 0, nil
}

func (s *Static) Update(todo *schema.Todo) error {
	return nil
}

func (s *Static) Delete(id int) error {
	return nil
}

func (s *Static) Close() {}
