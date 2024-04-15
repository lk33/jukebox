package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	ERR_MSG = "ERROR"
	MSG     = "SUCCESS"
)

func writeJSONMessage(w http.ResponseWriter, r *http.Request, msg string, msgType string, statusCode int) {
	data := jsonifyMessage(msg, msgType, statusCode)
	writeJSONResponse(w, r, data, statusCode)
}

func writeJSONStruct(w http.ResponseWriter, r *http.Request, v interface{}, statusCode int) {
	d, err := json.Marshal(v)
	if err != nil {
		writeJSONMessage(w, r, err.Error(), ERR_MSG, http.StatusInternalServerError)
		return
	}
	writeJSONResponse(w, r, d, statusCode)
}

func writeJSONResponse(w http.ResponseWriter, r *http.Request, d []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(d)
}

func jsonifyMessage(msg string, msgType string, statusCode int) []byte {

	var data []byte
	var obj struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	obj.Message = msg
	switch msgType {
	case ERR_MSG:
		obj.Status = "Failure"
	case MSG:
		obj.Status = "Success"
	}
	data, _ = json.Marshal(obj)
	return data
}
