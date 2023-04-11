package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectUri struct {
	PId int64 `uri:"p_id" binding:"required"`
}

type ProjectCreateUpdate struct {
	PName string `json:"p_name" binding:"required,lte=256"`
}

type TaskUri struct {
	TId int64 `uri:"t_id" binding:"required"`
}

type NewTask struct {
	TSubject string `json:"t_subject" binding:"required,lte=256"`
}

type Err struct {
	Error string `json:"error"`
}

func AbortWithBadUri(ctx *gin.Context, err error) {
	AbortWithBadRequest(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
}

func AbortWithBadRequest(ctx *gin.Context, message string) {
	AbortWithError(ctx, http.StatusBadRequest, message)
}

func AbortWithBadJson(ctx *gin.Context, err error) {
	AbortWithBadRequest(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
}

func AbortWithNotFound(ctx *gin.Context, message string) {
	AbortWithError(ctx, http.StatusNotFound, message)
}

func AbortWith500(ctx *gin.Context, message string) {
	AbortWithError(ctx, http.StatusInternalServerError, message)
}

func AbortWithError(ctx *gin.Context, httpStatusCode int, message string) {
	err := Err{
		Error: message,
	}
	AbortWithJSON(ctx, httpStatusCode, err)
}

func AbortWithJSON(ctx *gin.Context, httpStatusCode int, jsonObject interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(httpStatusCode, jsonObject)
}

func RespondWithJSON(ctx *gin.Context, httpStatusCode int, jsonObject interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(httpStatusCode, jsonObject)
}
