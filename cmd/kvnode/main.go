package main

import (
	"distributed-kv-store/internal/cluster"
	"distributed-kv-store/internal/config"
	"distributed-kv-store/internal/storage"
	th "distributed-kv-store/internal/transport/http"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	nodeID := flag.String("id", "node-1", "Unique ID for this node")
	port := flag.Int("port", 8001, "HTTP port to listen on")
	configPath := flag.String("config", "config/cluster.json", "path to cluster config")
	flag.Parse()

	cfg, err := config.LoadClusterConfig(*configPath)
	if err != nil {
		log.Fatalf("[%s] Failed to load cluster config: %v", *nodeID, err)
	}

	var peers []*cluster.Peer
	for _, p := range cfg.Peers{
		if p.ID == *nodeID{
			continue
		}
		peers = append(peers, cluster.NewPeer(p.ID, p.Addr))
	}

	healthChecker := cluster.NewHealthChecker(*nodeID, peers)
	healthChecker.Start(time.Second * 3)

	// storage and http server
	store := storage.NewStore()
	server := th.NewServer(*nodeID, store)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("[%s] Starting KV node on http://localhost%s", *nodeID, addr)

	err = http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatalf("[%s] Server failed: %v", *nodeID, err)
	}

}
