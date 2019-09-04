package main

import (
	"fmt"
	"time"
)

// Flow has a list of request
type Flow struct {
	requests        []*Request
	requestQueue    chan *Request
	addRequest      chan *Request
	pollingInterval time.Duration
}

// NewFlow is a factory method used to create an instance of the Flow struct.
// It starts the polling of the requests.
func NewFlow(requests ...*Request) Flow {
	flow := Flow{
		requests:        requests,
		requestQueue:    make(chan *Request),
		addRequest:      make(chan *Request),
		pollingInterval: minimumPollingInterval(requests),
	}

	flow.PollRequests()

	return flow
}

// PollRequests polls the requests of the Flow.
func (f *Flow) PollRequests() {
	for _, r := range f.requests {
		go func(r *Request) {
			r.Poll(f.requestQueue)
		}(r)
	}
}

// Go starts a goroutine that listens for requests in the queue.
// It's always ready to add new ones to its list.
func (f *Flow) Go() {
	go func() {
		for {
			select {

			case request := <-f.addRequest:
				f.requests = append(f.requests, request)
				f.pollingInterval = minimumPollingInterval(f.requests)
				request.Poll(f.requestQueue)

			case request := <-f.requestQueue:
				fmt.Println(time.Now().Format(time.StampMilli), request.GetMessage())
				<-time.After(f.pollingInterval)

			}
		}
	}()
}

// AddRequest adds a request to the list of the Flow's requests.
func (f *Flow) AddRequest(r *Request) {
	go func() {
		f.addRequest <- r
	}()
}
