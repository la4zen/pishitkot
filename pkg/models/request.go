package models

type Request struct {
	Code   string `form:"code" json:"code"`
	TaskID uint   `form:"task_id" json:"task_id"`
}
