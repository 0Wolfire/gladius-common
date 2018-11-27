package handlers

import (
	"net/http"
)

// VersionHandler - version information of the current module
func VersionHandler(version string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := make(map[string]string)
		result["version"] = version
		ResponseHandler(w, r, "Got version", true, nil, result, nil)
	}
}
