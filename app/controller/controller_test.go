package controller

import (
	"github.com/globalsign/mgo/bson"
	"github.com/wilsonfv/todolist/app/dao"
	"github.com/wilsonfv/todolist/app/model"
	"testing"
	"time"
)

func createMockDao() dao.MockTaskDao {
	var dao = dao.MockTaskDao{}

	creationDate1, _ := bson.NewMongoTimestamp(time.Now(), 1)
	creationDate2, _ := bson.NewMongoTimestamp(time.Now(), 2)
	creationDate3, _ := bson.NewMongoTimestamp(time.Now(), 3)

	var task1 = model.Task{
		ID:           bson.NewObjectId(),
		Name:         "task1",
		CreationDate: creationDate1,
		Description:  "task description 1"}

	var task2 = model.Task{
		ID:           bson.NewObjectId(),
		Name:         "task2",
		CreationDate: creationDate2,
		Description:  "task description 2"}

	var task3 = model.Task{
		ID:           bson.NewObjectId(),
		Name:         "task3",
		CreationDate: creationDate3,
		Description:  "task description 3"}

	dao.Collection = append(dao.Collection, task1)
	dao.Collection = append(dao.Collection, task2)
	dao.Collection = append(dao.Collection, task3)

	return dao
}

func TestListAll(t *testing.T) {
	var dao = createMockDao()

	tasks, _ := ListAll(&dao)

	if len(tasks) != 3 {
		t.Fail()
	}
}

func TestAddTask(t *testing.T) {
	var dao = createMockDao()

	creationDate4, _ := bson.NewMongoTimestamp(time.Now(), 4)
	var task4 = model.Task{
		ID:           bson.NewObjectId(),
		Name:         "task4",
		CreationDate: creationDate4,
		Description:  "task description 4"}

	AddTask(&dao, task4)

	tasks, _ := ListAll(&dao)

	if len(tasks) != 4 {
		t.Fail()
	}
}

func TestDeleteTask(t *testing.T) {
	var dao = createMockDao()

	taskListOld, _ := ListAll(&dao)

	var task = taskListOld[0]

	DeleteTask(&dao, task)

	taskListNew, _ := ListAll(&dao)

	if len(taskListNew) != 2 {
		t.Fail()
	}
}
