# sdm_demo_todolist_golang
Quick Demo of how to use [SQL DAL Maker](https://github.com/panedrone/sqldalmaker) + Go.

There are two parts:

* using Gorm
* using "database/sql" directly

Front-end is written in Vue.js, SQLite3 is used as a database.

![demo-go.png](demo-go.png)

dto.xml
```xml
<dto-class name="gorm-Project" ref="projects"/>

<dto-class name="ProjectLi" ref="get_projects.sql">

    <header>// Project list item</header>
    
    <field type="%int64" column="p_id"/>
    <field type="%string" column="p_name"/>
    <field type="%int64" column="p_tasks_count"/>

</dto-class>

<dto-class name="gorm-Task" ref="tasks"/>

<dto-class name="gorm-TaskLi" ref="tasks">

    <header>// Task list item (no p_id, no t_comments)</header>
    
    <field column="p_id" type="-"/>
    <field column="t_comments" type="-"/>

</dto-class>
```
ProjectsDao.xml
```xml
<crud dto="gorm-Project"/>

<query-dto-list method="ReadProjectList" dto="ProjectLi"/>
```
TasksDao.xml
```xml
<crud dto="gorm-Task"/>
```
Generated code in action:
```go
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
```