package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sdm_demo_todolist/gorm/dbal"
	"sdm_demo_todolist/gorm/models"
	"sdm_demo_todolist/util"
)

func ProjectCreate(ctx *gin.Context) {
	var inGr util.ProjectCreateUpdate
	err := ctx.ShouldBindJSON(&inGr)
	if err != nil {
		util.AbortWithBadJson(ctx, err)
		return
	}
	gr := models.Project{}
	gr.PName = inGr.PName
	err = dbal.NewProjectsDao().CreateProject(ctx, &gr)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func ProjectsReadAll(ctx *gin.Context) {
	groups, err := dbal.NewProjectsDao().ReadProjectList(ctx)
	if err != nil {
		util.AbortWith500(ctx, err)
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
	gr, err := dbal.NewProjectsDao().ReadProject(ctx, uri.PId)
	if err == gorm.ErrRecordNotFound {
		util.AbortWithNotFound(ctx, err.Error())
		return
	}
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
	util.RespondWithJSON(ctx, http.StatusOK, gr)
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
	dao := dbal.NewProjectsDao()
	pr, err := dao.ReadProject(ctx, uri.PId)
	if err != nil {
		util.AbortWithBadRequest(ctx, err.Error())
		return
	}
	pr.PName = inProject.PName
	_, err = dao.UpdateProject(ctx, pr)
	if err != nil {
		util.AbortWith500(ctx, err)
		return
	}
}

func ProjectDelete(ctx *gin.Context) {
	var uri util.ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		util.AbortWithBadUri(ctx, err)
		return
	}
	_, err := dbal.NewProjectsDao().DeleteProject(ctx, &models.Project{PId: uri.PId})
	if err != nil {
		util.AbortWith500(ctx, err)
	}
	ctx.Status(http.StatusNoContent)
}
