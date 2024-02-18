package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/andrei0427/lifeofmarrow/internal"
)

func HandlePurge(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("STRAPI_CACHE_KEY")
	keyHdr := r.Header.Get("X-API-Key")
	if key != keyHdr {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body to purge cache: %v", err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	strapiReq := struct {
		Model string `json:"model"`
	}{}
	err = json.Unmarshal(body, &strapiReq)
	if err != nil {
		log.Printf("error reading model info from strapi response: %s", err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	if _, ok := internal.Cache.Get(strapiReq.Model); ok {
		internal.Cache.Invalidate(strapiReq.Model)
	}

	w.WriteHeader(http.StatusOK)
}
