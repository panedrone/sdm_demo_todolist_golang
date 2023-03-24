# sdm_demo_todolist_sqlite3_golang
Quick Demo of how to use [SQL DAL Maker](https://github.com/panedrone/sqldalmaker) + Go: Gorm or direct "database/sql".

Both parts are sharing the same Vue.js front-end and the same SQLite3 database.

![demo-go.png](demo-go.png)

dto.xml
```xml
<dto-class name="gorm-Project" ref="projects"/>

<dto-class name="ProjectLi" ref="get_projects.sql">

    <header>// Group list item</header>
    
    <field type="int64${json}" column="p_id"/>
    <field type="string${json}" column="p_name"/>
    <field type="int64${json}" column="p_tasks_count"/>

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
<crud dto="gorm-Project" table="groups"/>

<query-dto-list method="GetAllProjects" dto="ProjectLi"/>
```
TasksDao.xml
```xml
<crud dto="gorm-Task" table="tasks"/>
```
Generated code in action:
```go
dao := dbal.NewProjectsDao()
// ---
err := dao.CreateProject(ctx, &project)
// ---
groups, err = dao.GetAllProjects(ctx)
// ---
group, err = dao.ReadProject(ctx, uri.PId)
// ---
ra, err = dao.UpdateProject(ctx, project)
// ---
ra, err = dao.DeleteProject(ctx, &models.Project{PId: uri.PId})
```