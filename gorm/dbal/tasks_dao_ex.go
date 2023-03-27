package dbal

import (
	"context"
	"sdm_demo_todolist/gorm/models"
)

// A hand-coded method extending functionality of generated class TasksDao

func (dao *TasksDao) ReadProjectTasks(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	err = dao.ds.Session(ctx).Table("tasks").Select("t_id", "t_date", "t_subject", "t_priority").
		Where("p_id = ?", pId).Order("t_date, t_id").
		Find(&res).Error
	return
}
