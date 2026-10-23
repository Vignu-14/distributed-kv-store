package http

import (
	"distributed-kv-store/internal/storage"
	"encoding/json"
	"net/http"
)


type Server struct{
	nodeId string
	store *storage.Store
	mux *http.ServeMux
}

func NewServer(nodeId string, store *storage.Store) *Server{
	s := &Server{
		nodeId: nodeId,
		store: store,
		mux: http.NewServeMux(),
	}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes(){
	s.mux.HandleFunc("kv", s.handleKV)
	s.mux.HandleFunc("internal/health", s.handleHealth)
}

func(s *Server) ServeHTTP(w  http.ResponseWriter, r *http.Request){
	s.mux.ServeHTTP(w, r)
}

func(s *Server)handleHealth(w http.ResponseWriter, r *http.Request){
	writeJSON(w, http.StatusOK, KVResponse{
		NodeId: s.nodeId,
		Status: "alive",
	})
}

func(s *Server) handleKV(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodPost:
		var req PutRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		s.store.Put(req.Key, req.Value)
		writeJSON(w, http.StatusCreated, KVResponse{
			NodeId: s.nodeId,
			Key: req.Key,
			Value: req.Value,
			Status: "stored",
		})
	
	case http.MethodGet:
		key := r.URL.Query().Get("key")
		if key == ""{
			http.Error(w, "invalid ?key= query parameter", http.StatusBadRequest)
			return
		}
		value, exists := s.store.Get(key)
		if !exists{
			writeJSON(w, http.StatusNotFound, KVResponse{
				NodeId: s.nodeId,
				Key: key,
				Status: "not_found",
			})
			return
		}
		writeJSON(w, http.StatusOK, KVResponse{
			NodeId: s.nodeId,
			Key: key,
			Value: value,
			Status: "ok",
		})
	case http.MethodDelete:
		key := r.URL.Query().Get("key")
		if key == ""{
			http.Error(w, "invalid ?key= query parameter", http.StatusBadRequest)
			return
		}
		s.store.Delete(key)
		writeJSON(w, http.StatusNoContent, KVResponse{
			NodeId: s.nodeId,
			Key: key,
			Status: "deleted",
		})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

