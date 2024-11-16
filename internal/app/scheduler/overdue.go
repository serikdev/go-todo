package scheduler

import (
	"time"

	"github.com/serikdev/go-todo/internal/app/model"
	"github.com/serikdev/go-todo/internal/app/repository"
)

func OverdueUpdater(stop chan struct{}) {
	for {
		select {
		case <-stop:
			return
		case <-time.After(1 * time.Minute):
			var tasks []model.Task
			repository.DB.Find(&tasks)

			for i := range tasks {
				if tasks[i].DueData != "" {
					if isOverdue(tasks[i].DueData) {
						tasks[i].OverData = true
						repository.DB.Save(&tasks[i])
					}
				}
			}
		}
	}
}

func isOverdue(dueDate string) bool {
	dueDateTime, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		return false
	}
	return time.Now().After(dueDateTime)
}
