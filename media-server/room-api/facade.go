package main

import (
	"github.com/google/uuid"
)

func createRoom() (*Room, *ServiceError) {
	rId := uuid.New().String()
	sId, herr := h.addRoom(rId)
	if herr != nil {
		return nil, herr
	}
	_, rerr := authoriseRoom(rId, sId)
	if rerr != nil {
		h.deleteRoom(rId)
		return nil, &ServiceError{
			StatusCode: RemoteCallFailed,
			Err:        rerr,
		}
	}
	return &Room{rId, sId}, nil
}

func joinRoom(rId roomId) (*Room, *ServiceError) {
	sid, herr := h.addTenant(rId)
	if herr != nil {
		return nil, herr
	}
	return &Room{rId, sid}, nil
}

func deleteRoom(rId roomId) *ServiceError {
	_, rerr := deauthoriseRoom(rId)
	if rerr != nil {
		return &ServiceError{
			StatusCode: RemoteCallFailed,
			Err:        rerr,
		}
	}
	h.deleteRoom(rId)
	return nil
}
