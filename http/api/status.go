package api

import "net/http"

func statusHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
}
