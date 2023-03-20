package dbal

// Code generated by a tool. DO NOT EDIT.
// Additional custom methods can be implemented in a separate file like <this_file>_ex.go.
// https://sqldalmaker.sourceforge.net/

import (
	"context"
	"sdm_demo_todolist/gorm/models"
)

type GroupsDao struct {
	ds DataStore
}

// (C)RUD: groups
// Generated/AI values are passed to DTO/model.

func (dao *GroupsDao) CreateGroup(ctx context.Context, p *models.Group) (err error) {
	err = dao.ds.Create(ctx, "groups", p)
	return
}

// C(R)UD: groups

func (dao *GroupsDao) ReadGroup(ctx context.Context, gId int64) (res *models.Group, err error) {
	res = &models.Group{}
	err = dao.ds.Read(ctx, "groups", res, gId)
	return
}

// CR(U)D: groups

func (dao *GroupsDao) UpdateGroup(ctx context.Context, p *models.Group) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Update(ctx, "groups", p)
	return
}

// CRU(D): groups

func (dao *GroupsDao) DeleteGroup(ctx context.Context, p *models.Group) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Delete(ctx, "groups", p)
	return
}

func (dao *GroupsDao) GetAllGroups(ctx context.Context) (res []*models.GroupLi, err error) {
	sql := `select g.*,  
		(select count(*) from tasks where g_id=g.g_id) as g_tasks_count 
		from groups g`
	err = dao.ds.QueryByFA(ctx, sql, &res)
	return
}

func (dao *GroupsDao) GetGroupLi(ctx context.Context, gId int64) (res *models.GroupLi, err error) {
	sql := `select g.*,  
		(select count(*) from tasks where g_id=g.g_id) as tasks_count 
		from groups g 
		where g_id=?`
	res = &models.GroupLi{}
	_fa := res
	err = dao.ds.QueryByFA(ctx, sql, _fa, gId)
	return
}

func (dao *GroupsDao) GetGroupLIIds(ctx context.Context) (res []int64, err error) {
	sql := `select g.*,  
		(select count(*) from tasks where g_id=g.g_id) as g_tasks_count 
		from groups g`
	errMap := make(map[string]int)
	onRow := func(val interface{}) {
		var data int64
		fromVal(&data, val, errMap)
		res = append(res, data)
	}
	err = dao.ds.QueryAll(ctx, sql, onRow)
	if err == nil {
		err = errMapToErr(errMap)
	}
	return
}

func (dao *GroupsDao) GetGroupLIId(ctx context.Context, gId string) (res int64, err error) {
	sql := `select g.*,  
		(select count(*) from tasks where g_id=g.g_id) as tasks_count 
		from groups g 
		where g_id=?`
	r, err := dao.ds.Query(ctx, sql, gId)
	if err == nil {
		err = assign(&res, r)
	}
	return
}

func (dao *GroupsDao) DeleteGroup2(ctx context.Context, gId string) (rowsAffected int64, err error) {
	sql := `delete from groups where g_id=?`
	rowsAffected, err = dao.ds.Exec(ctx, sql, gId)
	return
}
