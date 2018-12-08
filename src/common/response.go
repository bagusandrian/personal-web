package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONRespHeader struct {
	ProcessTime float64  `json:"process_time"`
	Messages    []string `json:"messages"`   // any message to be shown to user / client / caller
	Reason      string   `json:"reason"`     // Detailed cause why request failed
	ErrorCode   int64    `json:"error_code"` // Application Error Code

}

type JSONResponse struct {
	Header      JSONRespHeader `json:"header"`
	Data        interface{}    `json:"data"`
	StatusCode  int            `json:"-"` // HTTP Status Code
	ErrorString string         `json:"error,omitempty"`
	// TODO remove this field as we already have it in header
	Message string `jsong:"message"` // for backward compability with current codes in staging
	Log     string `json:"-"`
}

func (r *JSONResponse) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// TODO put proper allowed origin to avoid easy cyber attack
	w.Header().Set("Access-Control-Allow-Origin", "*")

	encoded, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR][SendResponse]: %+v\n", err)
	}

	if r.StatusCode != http.StatusOK {
		w.WriteHeader(r.StatusCode)
	}
	w.Write(encoded)
}
