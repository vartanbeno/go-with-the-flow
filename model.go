package main

import (
	"fmt"
	"sync"
	"time"
)

// Request holds an id and message.
// It's meant to be sent in a request queue repeatedly at its polling interval.
type Request struct {
	id      string
	message string

	pollingInterval time.Duration
	*sync.Mutex
}

// NewRequest is a factory method to create an instance of the Request struct.
// It makes sure the assigned id is unique, by appending the current Unix time in nanoseconds.
// It also takes care of initializing the mutex, which can be easily forgotten.
func NewRequest(id, message string, pollingInterval time.Duration) Request {
	return Request{
		id:              fmt.Sprintf("%s-%d", id, time.Now().UnixNano()),
		message:         message,
		pollingInterval: pollingInterval,
		Mutex:           &sync.Mutex{},
	}
}

// GetID returns the request's id.
func (r *Request) GetID() string {
	r.Lock()
	defer r.Unlock()
	return r.id
}

// GetMessage returns the request's message.
func (r *Request) GetMessage() string {
	r.Lock()
	defer r.Unlock()
	return r.message
}

// GetPollingInterval returns the request's polling interval.
func (r *Request) GetPollingInterval() time.Duration {
	r.Lock()
	defer r.Unlock()
	return r.pollingInterval
}

// Poll repeatedly sends the request into the request queue.
func (r *Request) Poll(requestQueue chan<- *Request) {
	go func() {
		for {
			requestQueue <- r
			<-time.After(r.GetPollingInterval())
		}
	}()
}
