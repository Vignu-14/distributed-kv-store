package cluster

import (
	"net/http"
	"time"
)

type HealthChecker struct {
	selfID string
	peers  []*Peer
	client *http.Client
}

func NewHealthChecker(selfId string, peers []*Peer) *HealthChecker{
	return &HealthChecker{
		selfID: selfId,
		peers: peers,
		client: &http.Client{
			Timeout: 1 * time.Second,
		},
	}
}


