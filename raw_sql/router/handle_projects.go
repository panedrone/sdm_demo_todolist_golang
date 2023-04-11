package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/raw_sql/dbal"
	"sdm_demo_todolist/raw_sql/dbal/dto"
	"sdm_demo_todolist/util"
)

var gDao = dbal.NewProjectsDao()

func ProjectCreate(ctx *gin.Context) {
	var inPr util.ProjectCreateUpdate
	err := ctx.ShouldBindJSON(&inPr)
	if err != nil {
		util.AbortWithBadJson(ctx, err)
		return
	}
	gr := dto.Project{}
	gr.PName = inPr.PName
	err = gDao.CreateProject(ctx, &gr)
	if err != nil {
		util.AbortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusCreated)
}

func ProjectsReadAll(ctx *gin.Context) {
	groups, err := gDao.GetProjectList(ctx)
	if err != nil {
		util.AbortWith500(ctx, err.Error())
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, groups)
}

func ProjectRead(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	group, err := gDao.ReadProject(ctx, uri.PId)
	if err != nil {
		if err == sql.ErrNoRows {
			util.AbortWithNotFound(ctx, err.Error())
		} else {
			util.AbortWith500(ctx, err.Error())
		}
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, group)
}

func ProjectUpdate(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	var inProject util.ProjectCreateUpdate
	err := ctx.ShouldBindJSON(&inProject)
	if err != nil {
		util.AbortWithBadJson(ctx, err)
		return
	}
	rowsAffected, err := gDao.UpdateProject(ctx, &dto.Project{PName: inProject.PName})
	if err != nil {
		util.AbortWith500(ctx, err.Error())
		return
	}
	if rowsAffected != 1 {
		util.AbortWith500(ctx, err.Error())
		return
	}
}

func ProjectDelete(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.BindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	rowsAffected, err := gDao.DeleteProject(ctx, &dto.Project{PId: uri.PId})
	if err != nil {
		util.AbortWith500(ctx, err.Error())
		return
	}
	if rowsAffected != 1 {
		util.AbortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusNoContent)
}
