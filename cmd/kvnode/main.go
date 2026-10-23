package main

import (
	"distributed-kv-store/internal/storage"
	th "distributed-kv-store/internal/transport/http"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	nodeID := flag.String("id", "node-1", "Unique ID for this node")
	port := flag.Int("port", 8001, "HTTP port to listen on")
	flag.Parse()

	store := storage.NewStore()
	server := th.NewServer(*nodeID, store)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("[%s] Starting KV node on http://localhost%s", *nodeID, addr)

	err := http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatalf("[%s] Server failed: %v", *nodeID, err)
	}

}
