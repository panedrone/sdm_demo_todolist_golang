package dbal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var ds = &_DS{}

func (ds *_DS) initDb() (err error) {
	ds.db, err = sql.Open("sqlite3", "./todo-list.sqlite")
	return
}

func OpenDB() error {
	return ds.Open()
}

func CloseDB() error {
	return ds.Close()
}

func NewTasksDao() *TasksDao {
	return &TasksDao{ds: ds}
}

func NewGroupsDao() *GroupsDao {
	return &GroupsDao{ds: ds}
}
