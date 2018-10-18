package main

import (
	"strconv"
	"fmt"
)

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	JobsQueue  chan Job

	maxWorkers int
	maxJobs    int
}

func NewDispatcher(maxJobsInQueue int, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, maxWorkers),
		JobsQueue:  make(chan Job, maxJobsInQueue),

		maxWorkers: maxWorkers,
		maxJobs: maxJobsInQueue,
	}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, strconv.Itoa(i))
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {

	for {
		select {
		case job := <-d.JobsQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool
				fmt.Println("[Dispatcher] there are", len(d.WorkerPool) + 1, "worker(s) available")
				// dispatch the job to the worker job channel
				jobChannel <- job
				fmt.Println("[Dispatcher] assign job", job, "to a worker")
			}(job)
		}
	}
}
