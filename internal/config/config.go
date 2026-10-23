package config

type PeerConfig struct{
	ID string `json:"id"`
	Addr string `json:"addr"`
}

type ClusterConfig struct{
	ReplicationFactor int `json:"replication_factor"`
	Peers []PeerConfig `json:"peers"`
}

