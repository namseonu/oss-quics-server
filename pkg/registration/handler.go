package registration

import (
	"encoding/json"
	"github.com/quic-s/quics/pkg/sync"
	"log"
	"net/http"

	"github.com/dgraph-io/badger/v3"
	"github.com/gorilla/mux"
	"github.com/quic-s/quics/pkg/client"
)

type Handler struct {
	db                  *badger.DB
	registrationService *Service
}

func NewRegistrationHandler(db *badger.DB) *Handler {
	clientRepository := client.NewClientRepository(db)
	registrationService := NewRegistrationService(clientRepository)
	return &Handler{db: db, registrationService: registrationService}
}

func (registrationHandler *Handler) SetupRoutes(r *mux.Router) {
	// 3.1 register root directory
	r.HandleFunc("/{clientId}/roots", registrationHandler.registerRootDir).Methods(http.MethodPost)

	// 3.2 get registered root directory
	r.HandleFunc("/{clientId}/roots/{rootId}", registrationHandler.getRegisteredRootDir).Methods(http.MethodGet)

	// 3.3 update reigstered root directory
	r.HandleFunc("/{clientId}/roots/{rootId}", registrationHandler.updateRegisteredRootDir).Methods(http.MethodPatch)

	// 3.4 unregister root directory
	r.HandleFunc("/{clientId}/roots/{rootId}", registrationHandler.deleteRegisteredRootDir).Methods(http.MethodPost)
}

func (registrationHandler *Handler) registerRootDir(w http.ResponseWriter, r *http.Request) {
	var request sync.RegisterRootDirRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Faield to create client because of wrong request data", http.StatusBadRequest)
		log.Fatalf("Error while decoding request data: %s", err)
		return
	}

	message, err := registrationHandler.registrationService.RegisterRootDir(request)
	if err != nil {
		http.Error(w, "Faield to register root directory", http.StatusInternalServerError)
		log.Fatalf("Error while returning response data: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func (registrationHandler *Handler) getRegisteredRootDir(w http.ResponseWriter, r *http.Request) {

}

func (registrationHandler *Handler) updateRegisteredRootDir(w http.ResponseWriter, r *http.Request) {

}

func (registrationHandler *Handler) deleteRegisteredRootDir(w http.ResponseWriter, r *http.Request) {

}
