package response

import (
	"encoding/json"
	"net/http"
)

func WriterJSON(values interface{}, w http.ResponseWriter) {
	dataResp, _ := json.MarshalIndent(values, "", " ")

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	_, err := w.Write(dataResp)
	if err != nil {
		return
	}
	_, _ = w.Write([]byte("\nInserted!!!!!\n"))
}
