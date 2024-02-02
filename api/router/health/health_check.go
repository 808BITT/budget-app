package health

import "net/http"

func Check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
