package node

import (
	"distributed-kv-store/internal/cluster"
	"distributed-kv-store/internal/storage"
	"net/http"
	"time"
)

type Node struct {
	ID         string
	Store      *storage.Store
	Peers      []*cluster.Peer
	httpClient http.Client
}

func NewNode(id string, store *storage.Store, peers []*cluster.Peer) *Node {
	return &Node{
		ID: id,
		Store: store,
		Peers: peers,
		httpClient: http.Client{
			Timeout: 2 * time.Second,
		},
	}
}
