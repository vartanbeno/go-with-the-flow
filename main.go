package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	request1 := NewRequest("id1", "this is a request", time.Millisecond*100)
	request2 := NewRequest("id2", "test request", time.Millisecond*100)

	flow := NewFlow(&request1, &request2)
	flow.Go()

	r := createRouter(&flow)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	// <-time.After(time.Second * 2)
}
