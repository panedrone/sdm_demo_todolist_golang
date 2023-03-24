package dbal

// Code generated by a tool. DO NOT EDIT.
// Additional custom methods can be implemented in a separate file like <this_file>_ex.go.
// https://sqldalmaker.sourceforge.net/

import (
	"context"
	"sdm_demo_todolist/raw_sql/dbal/dto"
)

type TasksDao struct {
	ds DataStore
}

// (C)RUD: tasks
// Generated/AI values are passed to DTO/model.

func (dao *TasksDao) CreateTask(ctx context.Context, p *dto.Task) (err error) {
	sql := `insert into tasks (p_id, t_priority, t_date, t_subject, t_comments) values (?, ?, ?, ?, ?)`
	res, err := dao.ds.Insert(ctx, sql, "t_id", p.PId, p.TPriority, p.TDate, p.TSubject, p.TComments)
	if err == nil {
		err = assign(&p.TId, res)
	}
	return
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTask(ctx context.Context, tId int64) (res *dto.Task, err error) {
	sql := `select * from tasks where t_id=?`
	res = &dto.Task{}
	_fa := []interface{}{
		&res.TId,
		&res.PId,
		&res.TPriority,
		&res.TDate,
		&res.TSubject,
		&res.TComments,
	}
	err = dao.ds.QueryByFA(ctx, sql, _fa, tId)
	return
}

// CR(U)D: tasks

func (dao *TasksDao) UpdateTask(ctx context.Context, p *dto.Task) (rowsAffected int64, err error) {
	sql := `update tasks set p_id=?, t_priority=?, t_date=?, t_subject=?, t_comments=? where t_id=?`
	rowsAffected, err = dao.ds.Exec(ctx, sql, p.PId, p.TPriority, p.TDate, p.TSubject, p.TComments, p.TId)
	return
}

// CRU(D): tasks

func (dao *TasksDao) DeleteTask(ctx context.Context, p *dto.Task) (rowsAffected int64, err error) {
	sql := `delete from tasks where t_id=?`
	rowsAffected, err = dao.ds.Exec(ctx, sql, p.TId)
	return
}

func (dao *TasksDao) GetGroupTasks(ctx context.Context, gId int64) (res []*dto.TaskLi, err error) {
	sql := `select t_id, t_priority, t_date, t_subject from tasks where p_id =? 
		order by t_id`
	_onRow := func() (interface{}, func()) {
		_obj := &dto.TaskLi{}
		return []interface{}{
				&_obj.TId,
				&_obj.TPriority,
				&_obj.TDate,
				&_obj.TSubject,
			}, func() {
				res = append(res, _obj)
			}
	}
	err = dao.ds.QueryAllByFA(ctx, sql, _onRow, gId)
	return
}

func (dao *TasksDao) DeleteGroupTasks(ctx context.Context, gId string) (rowsAffected int64, err error) {
	sql := `delete from tasks where p_id=?`
	rowsAffected, err = dao.ds.Exec(ctx, sql, gId)
	return
}

func (dao *TasksDao) GetCount(ctx context.Context) (res int64, err error) {
	sql := `select count(*) from tasks`
	r, err := dao.ds.Query(ctx, sql)
	if err == nil {
		err = assign(&res, r)
	}
	return
}

func (dao *TasksDao) GetGroupTasks2(ctx context.Context, gId int64) (res []*dto.TaskLi, err error) {
	sql := `select t_id, t_priority, t_date, t_subject from tasks where p_id=?`
	_onRow := func() (interface{}, func()) {
		_obj := &dto.TaskLi{}
		return []interface{}{
				&_obj.TId,
				&_obj.TPriority,
				&_obj.TDate,
				&_obj.TSubject,
			}, func() {
				res = append(res, _obj)
			}
	}
	err = dao.ds.QueryAllByFA(ctx, sql, _onRow, gId)
	return
}
