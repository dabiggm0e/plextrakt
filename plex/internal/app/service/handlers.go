package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/dabiggm0e/plextrakt/common/util"
	"github.com/go-chi/chi"
)

type Handler struct {
	//messagingClient messaging.IMessagingClient
	client    *http.Client
	myIP      string
	isHealthy bool
}

func NewHandler( /*messagingClient messaging.IMessageClient,*/ client *http.Client) *Handler {
	myIP, err := util.ResolveIPFromHostsFile()
	if err != nil {
		myIP = util.GetIP()
	}

	return &Handler{ /*messagingClient: messagingClient,*/ client: client, myIP: myIP, isHealthy: true}
}

// SetHealthyState can be used fake health problems.
func (h *Handler) SetHealthyState(w http.ResponseWriter, r *http.Request) {
	// Read the 'accountId' path parameter from the mux map
	var state, err = strconv.ParseBool(chi.URLParam(r, "state"))
	if err != nil {
		logrus.Errorln("Invalid request to SetHealthyState, allowed values are true or false")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.isHealthy = state
	w.WriteHeader(http.StatusOK)
}

func ParsePlexEvent(w http.ResponseWriter, r *http.Request) {
	//TODO: implement the handler
	data := []byte("Hai hai!")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Connection", "close")
	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
}

func (h *Handler) ParsePlexEvent(w http.ResponseWriter, r *http.Request) {
	//TODO: implement the handler
	data := []byte("Hai hai!")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Connection", "close")
	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
}

// HealthCheck will return OK if the underlying Messaging service is healthy.
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Since we're here, we already know that HTTP service is up. Let's just check the state of the boltdb connection
	// TODO: implement a health check for the messaging service
	messagingUp := true
	if messagingUp && h.isHealthy {
		data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
		writeJSONResponse(w, http.StatusOK, data)
	} else {
		data, _ := json.Marshal(healthCheckResponse{Status: "Messaging service is unaccessible"})
		writeJSONResponse(w, http.StatusServiceUnavailable, data)
	}
}

func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Connection", "close")
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}
