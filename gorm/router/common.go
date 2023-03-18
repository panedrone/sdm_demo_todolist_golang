package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type groupUri struct {
	GId int64 `uri:"g_id" binding:"required"`
}

type groupCreateUpdate struct {
	GName string `json:"g_name" binding:"required,lte=256"`
}

type taskUri struct {
	TId int64 `uri:"t_id" binding:"required"`
}

type newTask struct {
	TSubject string `json:"t_subject" binding:"required,lte=256"`
}

type Err struct {
	Error string `json:"error"`
}

func abortWithBadUri(ctx *gin.Context, err error) {
	abortWithBadRequest(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
}

func abortWithBadRequest(ctx *gin.Context, message string) {
	abortWithError(ctx, http.StatusBadRequest, message)
}

func abortWithNotFound(ctx *gin.Context, message string) {
	abortWithError(ctx, http.StatusNotFound, message)
}

func abortWith500(ctx *gin.Context, message string) {
	abortWithError(ctx, http.StatusInternalServerError, message)
}

func abortWithError(ctx *gin.Context, httpStatusCode int, message string) {
	err := Err{
		Error: message,
	}
	abortWithJSON(ctx, httpStatusCode, err)
}

func abortWithJSON(ctx *gin.Context, httpStatusCode int, jsonObject interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(httpStatusCode, jsonObject)
}

func respondWithJSON(ctx *gin.Context, httpStatusCode int, jsonObject interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(httpStatusCode, jsonObject)
}
