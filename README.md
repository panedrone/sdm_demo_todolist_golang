# sdm_demo_todolist_sqlite3_golang
Quick Demo of how to use [SQL DAL Maker](https://github.com/panedrone/sqldalmaker) + Go: Gorm and direct "database/sql".

Both parts are sharing the same Vue.js front-end and the same SQLite3 database.

![demo-go.png](demo-go.png)

![erd.png](erd.png)

dto.xml
```xml
<dto-class name="gorm-Group" ref="groups"/>

<dto-class name="GroupLi" ref="get_all_groups.sql">
    <header>// Group list item</header>
    <field type="int64${json}" column="g_id"/>
    <field type="string${json}" column="g_name"/>
    <field type="int64${json}" column="g_tasks_count"/>
</dto-class>

<dto-class name="gorm-Task" ref="tasks"/>

<dto-class name="gorm-TaskLi" ref="tasks">
    <header>// Task list item (no g_id, no t_comments)</header>
    <field column="g_id" type="-"/>
    <field column="t_comments" type="-"/>
</dto-class>
```
GroupsDao.xml
```xml
<crud dto="gorm-Group" table="groups"/>
<query-dto-list method="GetAllGroups" dto="GroupLi"/>
```
TasksDao.xml
```xml
<crud dto="gorm-Task" table="tasks"/>
```
Generated code in action:
```go
dao := dbal.NewGroupsDao()
// ---
err := dao.CreateGroup(ctx, &gr)
// ---
groups, err = dao.GetAllGroups(ctx)
// ---
group, err = dao.ReadGroup(ctx, uri.GId)
// ---
gr, err = dao.ReadGroup(ctx, uri.GId)
// ---
ra, err = dao.DeleteGroup(ctx, &models.Group{GId: uri.GId})
```