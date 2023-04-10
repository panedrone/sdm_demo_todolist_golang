package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sdm_demo_todolist/gorm/dbal"
	"sdm_demo_todolist/gorm/models"
	"time"
)

func TaskCreate(ctx *gin.Context) {
	var uri ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	var inTask NewTask
	err := ctx.ShouldBindJSON(&inTask)
	if err != nil {
		abortWithBadJson(ctx, err)
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
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusCreated)
}

func TaskRead(ctx *gin.Context) {
	var uri TaskUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	task, err := dbal.NewTasksDao().ReadTask(ctx, uri.TId)
	if err == gorm.ErrRecordNotFound {
		abortWithNotFound(ctx, err.Error())
	} else if err != nil {
		abortWith500(ctx, err.Error())
	} else {
		respondWithJSON(ctx, http.StatusOK, task)
	}
}

func TasksReadByProject(ctx *gin.Context) {
	var uri ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	var tasks []*models.TaskLi
	tasks, err := dbal.NewTasksDao().ReadProjectTasks(ctx, uri.PId)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	respondWithJSON(ctx, http.StatusOK, tasks)
}

func TaskUpdate(ctx *gin.Context) {
	var uri TaskUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	dao := dbal.NewTasksDao()
	t, err := dao.ReadTask(ctx, uri.TId)
	if err != nil {
		abortWithBadRequest(ctx, err.Error())
		return
	}
	var inTask models.Task
	err = ctx.ShouldBindJSON(&inTask)
	if err != nil {
		abortWithBadJson(ctx, err)
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
	t.TSubject = inTask.TSubject
	t.TPriority = inTask.TPriority
	t.TDate = inTask.TDate
	t.TComments = inTask.TComments
	_, err = dao.UpdateTask(ctx, t)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
}

func TaskDelete(ctx *gin.Context) {
	var uri TaskUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	_, err = dbal.NewTasksDao().DeleteTask(ctx, &models.Task{TId: uri.TId})
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusNoContent)
}
