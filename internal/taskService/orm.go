package taskservice

type RequestBodyTask struct {
	ID             int    `gorm:"primaryKey autoIncrement"`
	Task           string `json:"Task"`
	Accomplishment bool   `json:"Accomplishment"`
}

func (RequestBodyTask) TableName() string {
	return "request_body_tasks"
}
