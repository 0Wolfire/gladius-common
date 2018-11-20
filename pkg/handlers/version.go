package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// VersionHandler - version information of the current module
func VersionHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Open our jsonFile
		jsonFile, err := os.Open("version.json")
		if err != nil {
			fmt.Println(err)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result map[string]string
		json.Unmarshal([]byte(byteValue), &result)
		ResponseHandler(w, r, "Got version", true, nil, result, nil)
	}
}
