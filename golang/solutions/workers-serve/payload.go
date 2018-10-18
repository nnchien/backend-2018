package main

import (
	"time"
		"context"
	"math/rand"
	"fmt"
)

// Job represents the job to be run
type Job struct {
	Payload Payload
}

type Payload struct {
	Task string
}

type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

func (p *Payload) WorkOnAHeavyJob(workerTag string) error {
	time := randomTime()
	ctx, cancel := context.WithTimeout(context.Background(), time)
	defer cancel()

	select {
		case <-ctx.Done():
			fmt.Println(workerTag, *p, "in", time, "miliseconds")
			break
	}

	return nil
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func randomTime() time.Duration {
	return time.Duration(random(1, 10)* 100) *time.Millisecond
}
