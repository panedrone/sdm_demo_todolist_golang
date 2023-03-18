package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/raw_sql/dbal"
	"sdm_demo_todolist/raw_sql/dbal/dto"
)

var gDao = dbal.NewGroupsDao()

func groupCreate(ctx *gin.Context) {
	var inGr groupCreateUpdate
	err := ctx.ShouldBindJSON(&inGr)
	if err != nil {
		abortWithBadRequest(ctx, err.Error())
		return
	}
	gr := dto.Group{}
	gr.GName = inGr.GName
	err = gDao.CreateGroup(ctx, &gr)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusCreated)
}

func groupsReadAll(ctx *gin.Context) {
	groups, err := gDao.GetGroups(ctx)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	respondWithJSON(ctx, http.StatusOK, groups)
}

func groupRead(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	group, err := gDao.ReadGroup(ctx, uri.GId)
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

func groupUpdate(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.BindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	var inGr groupCreateUpdate
	err := ctx.ShouldBindJSON(&inGr)
	if err != nil {
		abortWithBadRequest(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	gr, err := gDao.ReadGroup(ctx, uri.GId)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	gr.GName = inGr.GName
	_, err = gDao.UpdateGroup(ctx, gr)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
}

func groupDelete(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.BindUri(&uri); err != nil {
		abortWithBadUri(ctx, err)
		return
	}
	gr := dto.Group{GId: uri.GId}
	_, err := gDao.DeleteGroup(ctx, &gr)
	if err != nil {
		abortWith500(ctx, err.Error())
		return
	}
	ctx.Status(http.StatusNoContent)
}
