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
		log.Printf("unauthorized request when trying to purge")
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

	if internal.Cache.Has(strapiReq.Model) {
		internal.Cache.Invalidate(strapiReq.Model)
	}

	// dirty hack, but works for now - if proves to be problematic, will just purge the entire thing
	if strapiReq.Model[len(strapiReq.Model)-1] != 's' {
		pluralKey := strapiReq.Model + "s"

		if internal.Cache.Has(pluralKey) {
			internal.Cache.Invalidate(pluralKey)
		}
	}

	w.WriteHeader(http.StatusOK)
}
