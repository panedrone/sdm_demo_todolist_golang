package dto

// Code generated by a tool. DO NOT EDIT.
// Additional custom methods can be implemented in a separate file like <this_file>_ex.go.
// https://sqldalmaker.sourceforge.net/

type Group struct {
	GId         int64  `json:"g_id"` // PK
	GName       string `json:"g_name"`
	GTasksCount int64  `json:"g_tasks_count"`
}
