package http

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, KVResponse{
		NodeId: s.node.ID,
		Status: "alive",
	})
}

func (s *Server) handleKV(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req PutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Key == "" {
			http.Error(w, "invalid request body: key and value required", http.StatusBadRequest)
			return
		}
		s.node.PutAndReplicate(req.Key, req.Value)
		writeJSON(w, http.StatusOK, KVResponse{
			NodeID: s.node.ID,
			Key:    req.Key,
			Value:  req.Value,
			Status: "stored_and_replicated",
		})
	case http.MethodGet:
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing ?key= query parameter", http.StatusBadRequest)
			return
		}
		val, exists := s.node.Store.Get(key)
		if !exists {
			writeJSON(w, http.StatusNotFound, KVResponse{
				NodeID: s.node.ID,
				Key:    key,
				Status: "not_found",
			})
			return
		}
		writeJSON(w, http.StatusOK, KVResponse{
			NodeID: s.node.ID,
			Key:    key,
			Value:  val,
			Status: "ok",
		})

	case http.MethodDelete:
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing ?key= query parameter", http.StatusBadRequest)
			return
		}
		s.node.DeleteAndReplicate(key)
		writeJSON(w, http.StatusOK, KVResponse{
			NodeID: s.node.ID,
			Key:    key,
			Status: "deleted_and_replicated",
		})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

	}
}

func (s *Server) handleReplicate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req PutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Key == "" {
			http.Error(w, "invalid replication payload", http.StatusBadRequest)
			return
		}
		// Save ONLY to local store (do NOT re-replicate!)
		s.node.Store.Put(req.Key, req.Value)
		writeJSON(w, http.StatusOK, KVResponse{
			NodeID: s.node.ID,
			Key:    req.Key,
			Status: "replicated",
		})

	case http.MethodDelete:
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing ?key= query parameter", http.StatusBadRequest)
			return
		}
		s.node.Store.Delete(key)
		writeJSON(w, http.StatusOK, KVResponse{
			NodeID: s.node.ID,
			Key:    key,
			Status: "replicated_delete",
		})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
