package router

import (
	"api/router/health"
	"api/util"
	"log"
	"net/http"
	"os"
)

func WebApp(w http.ResponseWriter, r *http.Request) {
	p := util.WebAssetPath() + r.URL.Path
	if _, err := os.Stat(p); os.IsNotExist(err) {
		http.ServeFile(w, r, util.WebAssetPath()+"/index.html")
	} else {
		http.ServeFile(w, r, p)
	}
}

func WebAssets(w http.ResponseWriter, r *http.Request) {
	assetPath := util.WebAssetPath() + r.URL.Path
	http.ServeFile(w, r, assetPath)
}

func RouterV1(w http.ResponseWriter, r *http.Request) {
	requestRoute := r.URL.Path[8:]
	log.Println(r.Method, requestRoute, r.RemoteAddr, r.UserAgent(), r.Header.Get("Content-Type"))
	switch requestRoute {
	case "health":
		health.Check(w, r)
	default:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	}
}
