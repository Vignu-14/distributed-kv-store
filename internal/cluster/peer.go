package cluster

import (
	"sync"
	"time"
)

type Peer struct {
	ID  string
	Addr string

	mu       sync.RWMutex
	alive    bool
	lastSeen time.Time
}

func NewPeer(id, addr string) *Peer {
	return &Peer{
		ID:    id,
		Addr:  addr,
		alive: false,
	}
}

// SetAlive updates the peer's status and returns true if the status changed
// (e.g., went from DEAD -> ALIVE or ALIVE -> DEAD).

// This funcion return true if the status of the Peer changes.
func (p *Peer) SetAlive(status bool) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	changed := p.alive != status
	p.alive = status
	if status {
		p.lastSeen = time.Now()
	}
	return changed
}

func(p *Peer) IsAlive()bool{
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.alive
}