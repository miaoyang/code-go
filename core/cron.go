package core

import "github.com/robfig/cron"

var Cron *TaskCron

type TaskCron struct {
	cron *cron.Cron
}

func InitCron() *TaskCron {
	Cron = &TaskCron{
		cron: cron.New(),
	}
	return Cron
}

// AddJob 添加定时任务
func (t *TaskCron) AddJob(jobName, spec string, job func()) error {
	LOG.Debug("开始执行定时任务：", jobName, " spec: ", spec)
	err := t.cron.AddFunc(spec, job)
	if err != nil {
		LOG.Error("添加任务失败")
		return err
	}
	return nil
}

func (t *TaskCron) Start() {
	t.cron.Start()
}

func (t *TaskCron) Stop() {
	t.cron.Stop()
}
