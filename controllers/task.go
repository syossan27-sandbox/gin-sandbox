package task

import (
  "gin-sandbox/models"
)

type Task struct {
}

func NewTask() Task {
  return Task{}
}

func (c Task) Get(n int) interface{} {
  repo := task.NewTaskRepository()
  task := repo.GetByID(n)
  return task
}

func (c Task) GetAll() interface{} {
  repo := task.NewTaskRepository()
  tasks := repo.GetAll()
  return tasks
}

func (c Task) Create(text string) {
  repo := task.NewTaskRepository()
  repo.Create(text)
}
