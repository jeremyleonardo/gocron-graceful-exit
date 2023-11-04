package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
)

func RunScheduler() {
	s := gocron.NewScheduler(time.UTC)

	s.WaitForScheduleAll() // to wait for the interval to pass before running the first task
	s.SingletonModeAll()   // a long running task will not be rescheduled until the current run is completed
	// FYI: if a task is running for longer than the the wait time, the next task can run immediately after the current task is completed

	_, err := s.Every(3).Second().Do(func() {
		TaskWrapper("task 1", func() {
			time.Sleep(5 * time.Second)
			// do something here
		})
	})
	if err != nil {
		log.Printf("error while scheduling task: %v", err)
	}

	_, err = s.Every(4).Second().Do(func() {
		TaskWrapper("task 2", func() {
			time.Sleep(5 * time.Second)
			// do something here
		})
	})
	if err != nil {
		log.Printf("error while scheduling task: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	go func() {
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	}()

	s.StartAsync()

	<-signalChan
	log.Println("received signal")
	s.Clear() // clear all the scheduled tasks
	log.Println("scheduler cleared")
	log.Println("stopping scheduler")

	// stop the scheduler, this will wait for all running tasks to complete first
	// there's no need for wg because the lib handles it internally
	s.Stop()

	log.Println("scheduler stopped")
	log.Println("might be waiting for running tasks to complete")
	log.Println("all tasks completed before exiting")
}

func TaskWrapper(taskName string, task func()) {
	log.Println(taskName + " - started")
	task()
	log.Println(taskName + " - finished")
}

func main() {
	RunScheduler()
}
