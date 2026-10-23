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


func (hc *HealthChecker) Start(interval time.Duration){
	ticker := time.NewTicker(interval)
	go func(){
		hc.checkAll()
		for range ticker.C{
			hc.checkAll()
		}
	}()
}