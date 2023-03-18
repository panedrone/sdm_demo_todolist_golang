package models

// Code generated by a tool. DO NOT EDIT.
// Additional custom methods can be implemented in a separate file like <this_file>_ex.go.
// https://sqldalmaker.sourceforge.net/

// TaskLI list item
type TaskLI struct {
	TId       int64  `json:"t_id" gorm:"column:t_id;primary_key;auto_increment"`
	TPriority int64  `json:"t_priority" gorm:"column:t_priority;not null"`
	TDate     string `json:"t_date" gorm:"column:t_date;not null"`
	TSubject  string `json:"t_subject" gorm:"column:t_subject;not null"`
}
