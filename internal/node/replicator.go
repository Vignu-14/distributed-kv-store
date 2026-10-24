package node

import (
	"bytes"
	"encoding/json"
	"log"
)

type ReplicatePayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (n *Node) replicatePutToPeers(key, value string) {
	body, err := json.Marshal(ReplicatePayload{Key: key, Value: value})
	if err != nil {
		log.Printf("[%s] Failed to marshal replication payload: %v", n.ID, err)
		return
	}

	for _, peer := range n.Peers{
		
		// check if peer is alive
		if !peer.IsAlive(){
			log.Printf("[%s] Skipping dead peer %s for key=%s", n.ID, peer.ID, key)
			continue
		}

		url := peer.Addr + "/internal/replicate"
		// passing json data by converting into io.Reader
		resp, err := n.httpClient.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			log.Printf("[%s] Failed to replicate key=%s to %s: %v", n.ID, key, peer.ID, err)
			continue
		}
		resp.Body.Close()	
		log.Printf("[%s] Replicated key=%s to peer %s", n.ID, key, peer.ID)
	}
}
