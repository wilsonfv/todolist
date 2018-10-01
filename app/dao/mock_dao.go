package dao

import (
	"errors"
	"github.com/wilsonfv/todolist/app/model"
)

type MockTaskDao struct {
	Collection []model.Task
}

func (td *MockTaskDao) Insert(task model.Task) error {
	if task.Name == "" {
		return errors.New("error")
	}

	if task.Description == "" {
		return errors.New("error")
	}

	td.Collection = append(td.Collection, task)
	return nil
}

func (td *MockTaskDao) Delete(task model.Task) error {
	var tasks []model.Task

	if task.ID == "" {
		return errors.New("error")
	}

	if task.Name == "" {
		return errors.New("error")
	}

	if task.Description == "" {
		return errors.New("error")
	}

	for _, aTask := range td.Collection {
		if aTask.ID != task.ID {
			tasks = append(tasks, aTask)
		}
	}

	td.Collection = tasks

	return nil
}

func (td *MockTaskDao) FindAll() ([]model.Task, error) {
	return td.Collection, nil
}
