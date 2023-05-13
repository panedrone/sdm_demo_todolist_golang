# sdm_demo_todolist_golang
Quick Demo of how to use [SQL DAL Maker](https://github.com/panedrone/sqldalmaker) + Go.

There are two parts:

* using Gorm
* using "database/sql" directly

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
dao := dbal.NewProjectsDao()
// ---
err := dao.CreateProject(ctx, &project)
// ---
projects, err = dao.ReadProjectList(ctx)
// ---
project, err = dao.ReadProject(ctx, uri.PId)
// ---
ra, err = dao.UpdateProject(ctx, project)
// ---
ra, err = dao.DeleteProject(ctx, &models.Project{PId: uri.PId})
```