package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/raw_sql/dbal"
	"sdm_demo_todolist/raw_sql/dbal/dto"
	"time"
)

var tDao = dbal.NewTasksDao()

func tasksReadByGroup(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	tasks, err := tDao.GetGroupTasks(ctx, uri.GId)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	respondWithJSON(ctx, http.StatusOK, tasks)
}

func taskRead(ctx *gin.Context) {
	var inTsk taskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		abortWithBadRequest(ctx, err.Error())
		return
	}
	task, err := tDao.ReadTask(ctx, inTsk.TId)
	if err == sql.ErrNoRows {
		abortWithNotFound(ctx, err.Error())
	} else if err != nil {
		abortWith500(ctx, err.Error())
	} else {
		respondWithJSON(ctx, http.StatusOK, task)
	}
}

func taskCreate(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	var inTask newTask
	err := ctx.ShouldBindJSON(&inTask)
	if err != nil {
		abortWithBadRequest(ctx, err.Error())
		return
	}
	t := dto.Task{}
	t.GId = uri.GId
	t.TSubject = inTask.TSubject
	t.TPriority = 1
	currentTime := time.Now().Local()
	layoutISO := currentTime.Format("2006-01-02")
	t.TDate = layoutISO
	err = tDao.CreateTask(ctx, &t)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusCreated)
}

func taskUpdate(ctx *gin.Context) {
	var inTsk taskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	var inTask dto.Task
	err = ctx.ShouldBindJSON(&inTask)
	if err != nil {
		abortWithBadRequest(ctx, err.Error())
		return
	}
	_, err = time.Parse("2006-01-02", inTask.TDate)
	if err != nil {
		abortWithBadRequest(ctx, fmt.Sprintf("Date format expected '2006-01-02': %s", err.Error()))
		return
	}
	if len(inTask.TSubject) == 0 {
		abortWithBadRequest(ctx, fmt.Sprintf("Subject required"))
		return
	}
	if inTask.TPriority <= 0 {
		abortWithBadRequest(ctx, fmt.Sprintf("Invalid Priority: %d", inTask.TPriority))
		return
	}
	t, err := tDao.ReadTask(ctx, inTsk.TId)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	t.TSubject = inTask.TSubject
	t.TPriority = inTask.TPriority
	t.TDate = inTask.TDate
	t.TComments = inTask.TComments
	_, err = tDao.UpdateTask(ctx, t)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
}

func taskDelete(ctx *gin.Context) {
	var inTsk taskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	t := dto.Task{TId: inTsk.TId}
	_, err = tDao.DeleteTask(ctx, &t)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusNoContent)
}
