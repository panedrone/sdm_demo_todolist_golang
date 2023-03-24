package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sdm_demo_todolist/gorm/dbal"
	"sdm_demo_todolist/gorm/models"
)

func ProjectCreate(ctx *gin.Context) {
	var inGr ProjectCreateUpdate
	err := ctx.ShouldBindJSON(&inGr)
	if err != nil {
		abortWithBadRequest(ctx, err.Error())
		return
	}
	gr := models.Project{}
	gr.PName = inGr.PName
	err = dbal.NewProjectsDao().CreateProject(ctx, &gr)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusCreated)
}

func ProjectsReadAll(ctx *gin.Context) {
	groups, err := dbal.NewProjectsDao().GetAllProjects(ctx)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	respondWithJSON(ctx, http.StatusOK, groups)
}

func ProjectRead(ctx *gin.Context) {
	var uri ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	gr, err := dbal.NewProjectsDao().ReadProject(ctx, uri.PId)
	if err == gorm.ErrRecordNotFound {
		abortWithNotFound(ctx, err.Error())
		return
	}
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	respondWithJSON(ctx, http.StatusOK, gr)
}

func ProjectUpdate(ctx *gin.Context) {
	var uri ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	var inProject ProjectCreateUpdate
	err := ctx.ShouldBindJSON(&inProject)
	if err != nil {
		abortWithBadRequest(ctx, err.Error())
		return
	}
	dao := dbal.NewProjectsDao()
	gr, err := dao.ReadProject(ctx, uri.PId)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	gr.PName = inProject.PName
	_, err = dao.UpdateProject(ctx, gr)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
}

func ProjectDelete(ctx *gin.Context) {
	var uri ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	_, err := dbal.NewProjectsDao().DeleteProject(ctx, &models.Project{PId: uri.PId})
	if err != nil {
		abortWith500(ctx, err.Error())
	}
	ctx.Status(http.StatusNoContent)
}
