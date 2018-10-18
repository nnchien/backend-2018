package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"time"
	"context"
)

var (
	MaxWorkers     = 2
	MaxJobsInQueue = 5

	// A buffered channel that we handle the workers on.
	DispatcherManager = NewDispatcher(MaxJobsInQueue, MaxWorkers)
)

func payloadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var collection PayloadCollection
	err = json.Unmarshal(b, &collection)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range collection.Payloads {

		// let's create a job with the payload
		work := Job{Payload: payload}

		// Push the work onto the queue.
		fmt.Println("[payloadHandler]", work, "-> JobsQueue")
		fmt.Println("[JobsQueue] has ", len(DispatcherManager.JobsQueue) + 1, "item(s)")
		DispatcherManager.JobsQueue <- work
	}

	output, err := json.Marshal(collection)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func statusTrigger() {
	jobsInQueue := 0
	workersInPool := 0
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		for {
			select {
				case <-ctx.Done():
					ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
					jobs := len(DispatcherManager.JobsQueue)
					workers := len(DispatcherManager.WorkerPool)
					if jobsInQueue != jobs || workersInPool != workers {
						fmt.Println("\n=== status updated ===")
						fmt.Println("JobsQueue has", len(DispatcherManager.JobsQueue), " job(s) in queue")
						fmt.Println("WorkerPool has", len(DispatcherManager.WorkerPool), " worker(s) in pool")
						fmt.Println("=== status updated ===\n")
						jobsInQueue = jobs
						workersInPool = workers
					}
					break
			}
		}
	}()
}

func main() {
	// start running dispatcher
	DispatcherManager.Run()

	// start status tracking
	statusTrigger()

	// start serve the endpoint
	http.HandleFunc("/", payloadHandler)
	http.ListenAndServe(":4000", nil)
}
