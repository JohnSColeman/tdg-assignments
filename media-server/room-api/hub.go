package main

import (
	"errors"
	"math"
)

// todo make configurable for production
const maxServerTenants = 300
const maxRoomCapacity = 300

type serverId = int
type roomId = string

type chatroomServer struct {
	id           int
	tenants      int
	roomsTenants map[roomId]int
}

func (s *chatroomServer) addTenant(rId roomId) {
	s.tenants++
	if s.roomsTenants == nil {
		s.roomsTenants = make(map[roomId]int)
	}
	s.roomsTenants[rId]++
}

type hub struct {
	servers map[serverId]*chatroomServer
}

var h = hub{
	servers: make(map[serverId]*chatroomServer),
}

func (h *hub) init(from int, count int) {
	for i := from; i < from+count; i++ {
		h.servers[i] = &chatroomServer{id: i, tenants: 0, roomsTenants: nil}
	}
}

func (h *hub) addRoom(rId roomId) (serverId serverId, serviceError *ServiceError) {
	if s := h.availableServer(); s != nil {
		s.addTenant(rId)
		return s.id, nil
	}
	return -1, &ServiceError{
		StatusCode: ServerNotAvailable,
		Err:        errors.New("no server available"),
	}
}

func (h *hub) addTenant(rId roomId) (serverId serverId, serviceError *ServiceError) {
	for sid, s := range h.servers {
		if tenants, found := s.roomsTenants[rId]; found {
			if tenants < maxRoomCapacity {
				s.addTenant(rId)
				return sid, nil
			}
			return -1, &ServiceError{
				StatusCode: RoomFull,
				Err:        errors.New("room is full"),
			}
		}
	}
	return -1, &ServiceError{
		StatusCode: RoomNotFound,
		Err:        errors.New("room not found"),
	}
}

func (h *hub) deleteRoom(rId roomId) {
	for _, s := range h.servers {
		if tenants, found := s.roomsTenants[rId]; found {
			s.tenants -= tenants
			delete(s.roomsTenants, rId)
		}
	}
}

func (h *hub) availableServer() *chatroomServer {
	h.cleanupRooms()
	var s *chatroomServer = nil
	var fs *chatroomServer = nil
	var minMessageRate float64 = math.MaxFloat64
	for _, cs := range h.servers {
		if cs.tenants < maxServerTenants && (s == nil || cs.tenants < s.tenants) {
			s = cs
		} else {
			smr, rerr := messageRate(string(cs.id))
			if rerr != nil && smr < minMessageRate {
				fs = cs
				minMessageRate = smr
			}
		}
	}
	if s != nil {
		return s
	}
	return fs
}

func (h *hub) cleanupRooms() {
	for _, rId := range closedRooms() {
		deauthoriseRoom(rId)
		h.deleteRoom(rId)
	}
}
