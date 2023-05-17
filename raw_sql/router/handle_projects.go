package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/raw_sql/dbal"
	"sdm_demo_todolist/raw_sql/dbal/dto"
	"sdm_demo_todolist/util"
)

var pDao = dbal.NewProjectsDao()

func ProjectCreate(ctx *gin.Context) {
	var inPr util.ProjectCreateUpdate
	err := ctx.ShouldBindJSON(&inPr)
	if err != nil {
		util.AbortWithBadJson(ctx, err)
		return
	}
	pr := dto.Project{}
	pr.PName = inPr.PName
	err = pDao.CreateProject(ctx, &pr)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func ProjectsReadAll(ctx *gin.Context) {
	pl, err := pDao.GetProjectList(ctx)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, pl)
}

func ProjectRead(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	pr, err := pDao.ReadProject(ctx, uri.PId)
	if err != nil {
		if err == sql.ErrNoRows {
			util.AbortWithNotFound(ctx, err.Error())
		} else {
			util.AbortWith500(ctx, err)
		}
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, pr)
}

func ProjectUpdate(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	var inPr util.ProjectCreateUpdate
	err := ctx.ShouldBindJSON(&inPr)
	if err != nil {
		util.AbortWithBadJson(ctx, err)
		return
	}
	_, err = pDao.UpdateProject(ctx, &dto.Project{PName: inPr.PName})
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
}

func ProjectDelete(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.BindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	_, err := pDao.DeleteProject(ctx, &dto.Project{PId: uri.PId})
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
