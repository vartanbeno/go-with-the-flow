package main

import "time"

func minimumPollingInterval(requests []*Request) time.Duration {
	minimum := requests[0]

	for _, r := range requests {
		if r.GetPollingInterval() < minimum.GetPollingInterval() {
			minimum = r
		}
	}

	return minimum.GetPollingInterval()
}
