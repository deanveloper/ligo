package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/deanveloper/ligo"
)

const homeUrl = "https://li.violet.wtf/"
const baseUrl = "https://lllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll.violet.wtf/"

func main() {
	http.HandleFunc("/", handleRedirect)

	// lol make this configurable/TLS if you want
	http.ListenAndServe(":8080", nil)
}

func handleRedirect(w http.ResponseWriter, req *http.Request) {
	code := strings.TrimPrefix(req.URL.Path, "/")
	if code == "" {
		http.Redirect(w, req, homeUrl, http.StatusPermanentRedirect)
		return
	}

	shortUrl, ok := ligo.CodeToWebsite(code)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest))
		return
	}
	http.Redirect(w, req, shortUrl, http.StatusPermanentRedirect)
}

func handleCreate(w http.ResponseWriter, req *http.Request) {
	link := req.URL.Query().Get("link")
	if link == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Fprint(w, baseUrl+ligo.WebsiteToCode(link, 200))
}
