package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/raw_sql/dbal"
	"sdm_demo_todolist/raw_sql/dbal/dto"
	"sdm_demo_todolist/util"
	"time"
)

var tDao = dbal.NewTasksDao()

func TaskCreate(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	var inTask util.NewTask
	err := ctx.ShouldBindJSON(&inTask)
	if err != nil {
		util.AbortWithBadRequest(ctx, err.Error())
		return
	}
	t := dto.Task{}
	t.PId = uri.PId
	t.TSubject = inTask.TSubject
	t.TPriority = 1
	currentTime := time.Now().Local()
	layoutISO := currentTime.Format("2006-01-02")
	t.TDate = layoutISO
	err = tDao.CreateTask(ctx, &t)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func TasksReadByProject(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	tasks, err := tDao.GetGroupTasks(ctx, uri.PId)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, tasks)
}

func TaskRead(ctx *gin.Context) {
	var inTsk util.TaskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		util.AbortWithBadRequest(ctx, err.Error())
		return
	}
	task, err := tDao.ReadTask(ctx, inTsk.TId)
	if err == sql.ErrNoRows {
		util.AbortWithNotFound(ctx, err.Error())
		return
	}
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, task)
}

func TaskUpdate(ctx *gin.Context) {
	var inTsk util.TaskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	var inTask dto.Task
	err = ctx.ShouldBindJSON(&inTask)
	if err != nil {
		util.AbortWithBadRequest(ctx, err.Error())
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
	t, err := tDao.ReadTask(ctx, inTsk.TId)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	t.TSubject = inTask.TSubject
	t.TPriority = inTask.TPriority
	t.TDate = inTask.TDate
	t.TComments = inTask.TComments
	_, err = tDao.UpdateTask(ctx, t)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
}

func TaskDelete(ctx *gin.Context) {
	var inTsk util.TaskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	t := dto.Task{TId: inTsk.TId}
	_, err = tDao.DeleteTask(ctx, &t)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
