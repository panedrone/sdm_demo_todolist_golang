package dto

// Code generated by a tool. DO NOT EDIT.
// Additional custom methods can be implemented in a separate file like <this_file>_ex.go.
// https://sqldalmaker.sourceforge.net/

type Project struct {
	PId         int64  `json:"p_id"`
	PName       string `json:"p_name"`
	PTasksCount int64  `json:"p_tasks_count"`
}
