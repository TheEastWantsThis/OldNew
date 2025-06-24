package taskservice

import (
	"gorm.io/gorm"
)

//ОСновные методв CRUD - Create Read Update Delete

type MainTaskRepository interface {
	CreateTask(task *RequestBodyTask) error
	GetAllTask() ([]RequestBodyTask, error)
	GetTaskByID(id int) (RequestBodyTask, error)
	UpdateTask(task RequestBodyTask) error
	DeleteTask(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) MainTaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task *RequestBodyTask) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) GetAllTask() ([]RequestBodyTask, error) {
	var task []RequestBodyTask
	err := r.db.Order("id asc").Find(&task).Error
	return task, err
}

func (r *taskRepository) GetTaskByID(id int) (RequestBodyTask, error) {
	var task RequestBodyTask
	err := r.db.First(&task, "ID = ?", id).Error
	return task, err
}

func (r *taskRepository) UpdateTask(task RequestBodyTask) error {
	return r.db.Save(&task).Error
}

func (r *taskRepository) DeleteTask(id int) error {
	return r.db.Delete(&RequestBodyTask{}, "id = ?", id).Error // Возможная проблема в id
}
