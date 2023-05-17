package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sdm_demo_todolist/gorm/dbal"
	"sdm_demo_todolist/gorm/models"
	"sdm_demo_todolist/util"
	"time"
)

func TaskCreate(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	var inTask util.NewTask
	err := ctx.ShouldBindJSON(&inTask)
	if err != nil {
		util.AbortWithBadJson(ctx, err)
		return
	}
	t := models.Task{}
	t.PId = uri.PId
	t.TSubject = inTask.TSubject
	t.TPriority = 1
	currentTime := time.Now().Local()
	layoutISO := currentTime.Format("2006-01-02")
	t.TDate = layoutISO
	err = dbal.NewTasksDao().CreateTask(ctx, &t)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func TaskRead(ctx *gin.Context) {
	var uri util.TaskUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	task, err := dbal.NewTasksDao().ReadTask(ctx, uri.TId)
	if err == gorm.ErrRecordNotFound {
		util.AbortWithNotFound(ctx, err.Error())
		return
	}
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, task)
}

func TasksReadByProject(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	var tasks []*models.TaskLi
	tasks, err := dbal.NewTasksDao().ReadProjectTasks(ctx, uri.PId)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, tasks)
}

func TaskUpdate(ctx *gin.Context) {
	var uri util.TaskUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	dao := dbal.NewTasksDao()
	t, err := dao.ReadTask(ctx, uri.TId)
	if err != nil {
		util.AbortWithBadRequest(ctx, err.Error())
		return
	}
	var inTask models.Task
	err = ctx.ShouldBindJSON(&inTask)
	if err != nil {
		util.AbortWithBadJson(ctx, err)
		return
	}
	_, err = time.Parse("2006-01-02", inTask.TDate)
	if err != nil {
		util.AbortWithBadRequest(ctx, fmt.Sprintf("Date format expected '2006-01-02': %s", err.Error()))
		return
	}
	if len(inTask.TSubject) == 0 {
		util.AbortWithBadRequest(ctx, fmt.Sprintf("Subject required"))
		return
	}
	if inTask.TPriority <= 0 {
		util.AbortWithBadRequest(ctx, fmt.Sprintf("Invalid Priority: %d", inTask.TPriority))
		return
	}
	t.TSubject = inTask.TSubject
	t.TPriority = inTask.TPriority
	t.TDate = inTask.TDate
	t.TComments = inTask.TComments
	_, err = dao.UpdateTask(ctx, t)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
}

func TaskDelete(ctx *gin.Context) {
	var uri util.TaskUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	_, err = dbal.NewTasksDao().DeleteTask(ctx, &models.Task{TId: uri.TId})
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
