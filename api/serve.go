package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/freneticmonkey/migrate/go/util"
	"github.com/gorilla/mux"
)

// ResponseError Standardised response error helper struct
type ResponseError struct {
	Error  string      `json:"error"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}

// Response Standardised response helper struct
type Response struct {
	Result interface{}   `json:"result"`
	Error  ResponseError `json:"error"`
}

// health return structure
type health struct {
	Version int    `json:"version"`
	Status  string `json:"status"`
	Context string `json:"context"`
}

// Config is used to pass through server configuration options
type Config struct {
	IP   string
	Port int
}

// writeResponse Helper function for building a standardised JSON response
func writeResponse(w http.ResponseWriter, body interface{}, e error) error {
	var (
		payload  []byte
		response Response
		err      error
	)
	response = Response{
		Result: body,
		Error:  ResponseError{},
	}
	if e != nil {
		response.Error.Error = fmt.Sprintf("%v", e)
	}
	payload, err = json.MarshalIndent(response, "", "\t")

	if !util.ErrorCheck(err) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
	return err
}

type endpoints struct {
	Version int
}

func (e endpoints) GetHealth(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, health{
		Version: e.Version,
		Status:  "OK",
		Context: "Normal",
	}, nil)
}

// Start the API server
func Start(cfg Config) {
	r := mux.NewRouter()

	ep := endpoints{
		Version: 1,
	}

	// Health Check
	r.HandleFunc("/api/health/", ep.GetHealth)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%d", cfg.IP, cfg.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
