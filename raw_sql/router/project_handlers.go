package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/raw_sql/dbal"
	"sdm_demo_todolist/raw_sql/dbal/dto"
)

var gDao = dbal.NewProjectsDao()

func ProjectCreate(ctx *gin.Context) {
	var inGr projectCreateUpdate
	err := ctx.ShouldBindJSON(&inGr)
	if err != nil {
		abortWithBadRequest(ctx, err.Error())
		return
	}
	gr := dto.Project{}
	gr.PName = inGr.PName
	err = gDao.CreateProject(ctx, &gr)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusCreated)
}

func ProjectsReadAll(ctx *gin.Context) {
	groups, err := gDao.GetProjects(ctx)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	respondWithJSON(ctx, http.StatusOK, groups)
}

func ProjectRead(ctx *gin.Context) {
	var uri projectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	group, err := gDao.ReadProject(ctx, uri.PId)
	if err != nil {
		if err == sql.ErrNoRows {
			abortWithNotFound(ctx, err.Error())
		} else {
			abortWith500(ctx, err.Error())
		}
		return
	}
	respondWithJSON(ctx, http.StatusOK, group)
}

func ProjectUpdate(ctx *gin.Context) {
	var uri projectUri
	if err := ctx.BindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	var inProject projectCreateUpdate
	err := ctx.ShouldBindJSON(&inProject)
	if err != nil {
		abortWithBadRequest(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	rowsAffected, err := gDao.UpdateProject(ctx, &dto.Project{PName: inProject.PName})
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	if rowsAffected != 1 {
		abortWith500(ctx, err.Error())
		return
	}
}

func ProjectDelete(ctx *gin.Context) {
	var uri projectUri
	if err := ctx.BindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	rowsAffected, err := gDao.DeleteProject(ctx, &dto.Project{PId: uri.PId})
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	if rowsAffected != 1 {
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusNoContent)
}
