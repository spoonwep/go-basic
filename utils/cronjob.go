package utils

import "github.com/robfig/cron/v3"

func InitCronjob() {
	c := cron.New(cron.WithSeconds())
	//Add cron jobs here

	//End cron jobs
	c.Start()
}
