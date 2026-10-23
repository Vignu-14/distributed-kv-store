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

func