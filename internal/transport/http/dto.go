package http

type PutRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type KVResponse struct {
	NodeId string `json:"node_id"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
	Status string `json:"status"`
}

