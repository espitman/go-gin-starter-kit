package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func TestTask() {
	fmt.Println("I am running TestTask.")
}

func (t *T) TestRun() {
	s1 := gocron.NewScheduler(time.UTC)
	_, _ = s1.Every(1).Seconds().Do(TestTask)
	s1.StartAsync()
}
