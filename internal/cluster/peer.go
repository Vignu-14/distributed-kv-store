package cluster

import (
	"sync"
	"time"
)

type Peer struct {
	Id   string
	Addr string

	mu       sync.RWMutex
	alive    bool
	lastSeen time.Time
}

func NewPeer(id, addr string) *Peer{
	return &Peer{
		Id: id,
		Addr: addr,
		alive: false,
	}
}

// SetAlive updates the peer's status and returns true if the status changed
// (e.g., went from DEAD -> ALIVE or ALIVE -> DEAD).
func