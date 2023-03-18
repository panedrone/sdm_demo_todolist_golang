package dbal

import (
	"context"
	"sdm_demo_todolist/gorm/models"
)

// A hand-coded method extending functionality of generated class TasksDao

func (dao *TasksDao) ReadGroupTasks(ctx context.Context, gId int64) (res []*models.TaskLI, err error) {
	// Don't fetch "t_comments" for Group-Tasks list:
	err = dao.ds.Session(ctx).Table("tasks").Where("g_id = ?", gId).Order("t_date, t_id").
		Select("t_id", "t_date", "t_subject", "t_priority").Find(&res).Error
	return
}
