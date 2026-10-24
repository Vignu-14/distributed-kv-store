package http

import "net/http"

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request){
	writeJSON(w, http.StatusOK, KVResponse{
		NodeId: s.node.ID,
		Status: "alive",
	})
}


