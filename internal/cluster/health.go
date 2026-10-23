package cluster

import (
	"log"
	"net/http"
	"time"
)

type HealthChecker struct {
	selfID string
	peers  []*Peer
	client *http.Client
}

func NewHealthChecker(selfId string, peers []*Peer) *HealthChecker {
	return &HealthChecker{
		selfID: selfId,
		peers:  peers,
		client: &http.Client{
			Timeout: 1 * time.Second,
		},
	}
}

func (hc *HealthChecker) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		hc.checkAll()
		for range ticker.C {
			hc.checkAll()
		}
	}()
}

func (hc *HealthChecker) checkAll() {
	for _, peer := range hc.peers {
		url := peer.Addr + "/internal/health"
		resp, err := hc.client.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			if resp != nil {
				resp.Body.Close()
			}
			if peer.SetAlive(false) {
				log.Printf("[%s] Peer %s (%s) is now DEAD", hc.selfID, peer.Id, peer.Addr)
			}
			continue
		}
		resp.Body.Close()

		if peer.SetAlive(true) {
			log.Printf("[%s] Peer %s (%s) is now ALIVE", hc.selfID, peer.Id, peer.Addr)
		}
	}
}
