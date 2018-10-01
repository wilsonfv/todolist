package controller

import (
	"github.com/globalsign/mgo/bson"
	"github.com/wilsonfv/todolist/app/dao"
	"github.com/wilsonfv/todolist/app/model"
	"time"
)

func ListAll(dao dao.Dao) ([]model.Task, error) {
	tasks, err := dao.FindAll()
	return tasks, err
}

func AddTask(dao dao.Dao, task model.Task) error {
	task.ID = bson.NewObjectId()

	if ts, err := bson.NewMongoTimestamp(time.Now(), 0); err != nil {
		return err
	} else {
		task.CreationDate = ts
	}

	if err := dao.Insert(task); err != nil {
		return err
	}

	return nil
}

func DeleteTask(dao dao.Dao, task model.Task) error {
	err := dao.Delete(task)
	return err
}

