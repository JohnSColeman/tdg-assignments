package main

import "fmt"

const (
	ServerNotAvailable = iota
	RoomNotFound       = iota
	ServerFull         = iota
	RoomFull           = iota
	RemoteCallFailed   = iota
)

type ServiceError struct {
	StatusCode int
	Err        error
}

func (r *ServiceError) Error() string {
	return fmt.Sprintf("status %d: %v", r.StatusCode, r.Err)
}
