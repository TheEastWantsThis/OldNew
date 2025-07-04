package taskservice

type MainTaskService interface {
	CreateTask(req RequestBodyTask) (RequestBodyTask, error) //Уточнить что писать в начале
	GetAllTask() ([]RequestBodyTask, error)
	GetTaskByID(id int) (RequestBodyTask, error)
	UpdateTask(id int, task RequestBodyTask) (RequestBodyTask, error)
	DeleteTask(id int) error
}

type taskService struct {
	repo MainTaskRepository
}

func NewTaskService(repo MainTaskRepository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) CreateTask(req RequestBodyTask) (RequestBodyTask, error) {

	postTask := RequestBodyTask{
		Task:           req.Task,
		Accomplishment: req.Accomplishment,
	}

	if err := s.repo.CreateTask(&postTask); err != nil {
		return RequestBodyTask{}, err
	}
	return postTask, nil
}

func (s *taskService) GetAllTask() ([]RequestBodyTask, error) {
	return s.repo.GetAllTask()
}

func (s *taskService) GetTaskByID(id int) (RequestBodyTask, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id int, task RequestBodyTask) (RequestBodyTask, error) {
	newtask, err := s.repo.GetTaskByID(id)
	if err != nil {
		return RequestBodyTask{}, err
	}

	newtask.Task = task.Task
	newtask.Accomplishment = task.Accomplishment

	if err := s.repo.UpdateTask(newtask); err != nil {
		return RequestBodyTask{}, err
	}
	return newtask, nil
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}
