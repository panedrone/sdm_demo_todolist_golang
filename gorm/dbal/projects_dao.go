package dbal

// Code generated by a tool. DO NOT EDIT.
// Additional custom methods can be implemented in a separate file like <this_file>_ex.go.
// https://sqldalmaker.sourceforge.net/

import (
	"context"
	"sdm_demo_todolist/gorm/models"
)

type ProjectsDao struct {
	ds DataStore
}

// (C)RUD: projects
// Generated/AI values are passed to DTO/model.

func (dao *ProjectsDao) CreateProject(ctx context.Context, p *models.Project) error {
	return dao.ds.Create(ctx, "projects", p)
}

// C(R)UD: projects

func (dao *ProjectsDao) ReadProject(ctx context.Context, pId int64) (*models.Project, error) {
	res := &models.Project{}
	err := dao.ds.Read(ctx, "projects", res, pId)
	if err == nil {
		return res, nil
	}
	return nil, err
}

// CR(U)D: projects

func (dao *ProjectsDao) UpdateProject(ctx context.Context, p *models.Project) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Update(ctx, "projects", p)
	return
}

// CRU(D): projects

func (dao *ProjectsDao) DeleteProject(ctx context.Context, p *models.Project) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Delete(ctx, "projects", p)
	return
}

func (dao *ProjectsDao) ReadProjectList(ctx context.Context) (res []*models.ProjectLi, err error) {
	sql := `select p.*, 
		(select count(*) from tasks where p_id=p.p_id) as p_tasks_count 
		from projects p 
		order by p.p_id`
	errMap := make(map[string]int)
	_onRow := func(row map[string]interface{}) {
		obj := models.ProjectLi{}
		SetInt64(&obj.PId, row, "p_id", errMap)
		SetString(&obj.PName, row, "p_name", errMap)
		SetInt64(&obj.PTasksCount, row, "p_tasks_count", errMap)
		res = append(res, &obj)
	}
	err = dao.ds.QueryAllRows(ctx, sql, _onRow)
	if err == nil {
		err = ErrMapToErr(errMap)
	}
	return
}

func (dao *ProjectsDao) GetProjectLIds(ctx context.Context) (res []int64, err error) {
	sql := `select p.*, 
		(select count(*) from tasks where p_id=p.p_id) as p_tasks_count 
		from projects p 
		order by p.p_id`
	errMap := make(map[string]int)
	onRow := func(val interface{}) {
		var data int64
		SetScalarValue(&data, val, errMap)
		res = append(res, data)
	}
	err = dao.ds.QueryAll(ctx, sql, onRow)
	if err == nil {
		err = ErrMapToErr(errMap)
	}
	return
}

func (dao *ProjectsDao) DeleteProject2(ctx context.Context, pId string) (rowsAffected int64, err error) {
	sql := `delete from projects where p_id=?`
	rowsAffected, err = dao.ds.Exec(ctx, sql, pId)
	return
}
