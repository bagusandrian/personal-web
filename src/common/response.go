package common

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
type HTMLResponse struct {
	Header      JSONRespHeader `json:"header"`
	FileHTML    string         `json:"file"`
	Data        interface{}    `json:"data"`
	StatusCode  int            `json:"-"` // HTTP Status Code
	ErrorString string         `json:"error,omitempty"`
	Log         string         `json:"-"`
}
type FileResponse struct {
	Path string `json:path"`
	Type string `json:"type"`
	File string `json:"file"`
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

func (r *HTMLResponse) SendResponse(w http.ResponseWriter) {
	now := time.Now()
	tmpl := template.Must(template.ParseFiles("files/var/www/index.html"))
	data := r.Data
	tmpl.Execute(w, data)
	log.Printf("process time %+v\n", time.Since(now))
}

func (r *FileResponse) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", r.Type)
	content := []byte{}
	var err error
	pathFolder := []string{
		"files/var/www/",
		"/files/var/www/",
		"./files/var/www/",
		"../files/var/www/",
		"../../files/var/www/",
		"../../../files/var/www/",
	}
	s := ""
	for _, p := range pathFolder {
		content, err = ioutil.ReadFile(p + r.Path + r.File)
		s = p
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Printf("error get file %s%s%s err: %+v\n", s, r.Path, r.File, err)
	}
	w.Write(content)
}
