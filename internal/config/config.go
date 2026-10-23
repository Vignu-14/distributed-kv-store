package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type PeerConfig struct{
	ID string `json:"id"`
	Addr string `json:"addr"`
}

type ClusterConfig struct{
	ReplicationFactor int `json:"replication_factor"`
	Peers []PeerConfig `json:"peers"`
}

func LoadClusterConfig(path string) (*ClusterConfig, error){
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	
	var cfg ClusterConfig
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &cfg, nil
}