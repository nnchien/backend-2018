package main

import "fmt"

const WORKER_LOG_FORMAT = "worker(%s) [%s] %s"

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
	name       string
}

func NewWorker(workerPool chan chan Job, name string) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		name:       name,
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			fmt.Println(fmt.Sprintf(WORKER_LOG_FORMAT, w.name, "Available", "register to worker queue"))

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				fmt.Println(fmt.Sprintf(WORKER_LOG_FORMAT, w.name, "Received Job", job))
				fmt.Println(fmt.Sprintf(WORKER_LOG_FORMAT, w.name, "Working On", "WorkOnAHeavyJob"), job)
				if err := job.Payload.WorkOnAHeavyJob(fmt.Sprintf(WORKER_LOG_FORMAT, w.name, "Done", "WorkOnAHeavyJob")); err != nil {
					// never happens
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}
