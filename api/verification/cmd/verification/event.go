package main

import "github.com/utilitywarehouse/partner-pkg/event"

func NewEvent(id string) event.Event {
	return event.Event{
		ID:      "123",
		Comment: "foo",
	}
}
