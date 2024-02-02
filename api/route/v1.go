package route

import (
	"api/util"
	"log"
	"net/http"
	"os"
)

func Router(w http.ResponseWriter, r *http.Request) {
	requestRoute := r.URL.Path[8:]
	log.Println(r.Method, requestRoute, r.RemoteAddr, r.UserAgent(), r.Header.Get("Content-Type"))
	switch requestRoute {
	case "health":
		health(w, r)
	default:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func WebApp(w http.ResponseWriter, r *http.Request) {
	p := util.AssetPath() + r.URL.Path
	if _, err := os.Stat(p); os.IsNotExist(err) {
		http.ServeFile(w, r, util.AssetPath()+"/index.html")
	} else {
		http.ServeFile(w, r, p)
	}
}

func ServeAsset(w http.ResponseWriter, r *http.Request) {
	assetPath := util.AssetPath() + r.URL.Path
	http.ServeFile(w, r, assetPath)
}
