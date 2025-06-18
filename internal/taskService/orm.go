package taskservice

type RequestBodyTask struct {
	ID             int    `gorm:"primaryKey autoIncrement"` //json:"ID"
	Task           string `json:"Task"`
	Accomplishment bool   `json:"Accomplishment"`
}
