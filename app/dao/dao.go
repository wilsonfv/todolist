package dao

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/wilsonfv/todolist/app/model"
	"log"
)

type Dao interface {
	Insert(task model.Task) error
	Delete(task model.Task) error
	FindAll() ([]model.Task, error)
}

type TaskDao struct {
	Server     string
	Database   string
	Collection string
	C          *mgo.Collection
}

func (td *TaskDao) Connect() {
	session, err := mgo.Dial(td.Server)

	if err != nil {
		log.Fatal(err)
	}

	td.C = session.DB(td.Database).C(td.Collection)
}

func (td *TaskDao) Insert(task model.Task) error {
	err := td.C.Insert(&task)
	return err
}

func (td *TaskDao) Delete(task model.Task) error {
	err := td.C.Remove(&task)
	return err
}

func (td *TaskDao) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	err := td.C.Find(bson.M{}).All(&tasks)
	return tasks, err
}
