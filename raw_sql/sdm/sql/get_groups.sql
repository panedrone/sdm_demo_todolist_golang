select g.*, 
(select count(*) from tasks where g_id=g.g_id) as g_tasks_count
from groups g