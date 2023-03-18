package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/gorm/dbal"
)

func New() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/
	// === panedrone: type "http://localhost:8080/assets/" to render index.html

	static := r.Static("/static", "./static")
	static.StaticFile("/favicon.ico", "./static/go.png")

	static.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/static")
	})

	api := r.Group("", func(ctx *gin.Context) { // middleware
		ctx.Set("db", dbal.WithContext(ctx))
	})

	r.GET("/whoiam", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Gorm")
	})
	{
		groups := api.Group("/groups")
		groups.GET("/", groupsReadAll)
		groups.POST("/", groupCreate)
		{
			group := groups.Group("/:g_id")
			group.GET("/", groupRead)
			group.PUT("/", groupUpdate)
			group.DELETE("/", groupDelete)
			{
				groupTasks := group.Group("/tasks")
				groupTasks.GET("/", tasksReadByGroup)
				groupTasks.POST("/", taskCreate)
			}
		}
	}
	{
		task := api.Group("/tasks/:t_id")
		task.GET("", taskRead)
		task.PUT("", taskUpdate)
		task.DELETE("", taskDelete)
	}

	return r
}
