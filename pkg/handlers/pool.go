package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gladiusio/gladius-cli/utils"
	"github.com/gladiusio/gladius-common/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
	"github.com/gorilla/mux"
)

// PoolPublicDataHandler gets the public data of a specified pool
func PoolPublicDataHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		vars := mux.Vars(r)
		poolAddress := vars["poolAddress"]

		poolResponse, err := PoolResponseForAddress(poolAddress, ga)

		if err != nil {
			ErrorHandler(w, r, "Pool data could not be found for Pool: "+poolAddress, err, http.StatusBadRequest)
			return
		}

		poolInformationResponse, err := utils.SendRequest(http.MethodGet, poolResponse.Url+"server/info", nil)
		var defaultResponse response.DefaultResponse
		json.Unmarshal([]byte(poolInformationResponse), &defaultResponse)

		ResponseHandler(w, r, "null", true, nil, defaultResponse.Response, nil)
	}
}

// MarketPoolsHandler - Returns all Pools
func MarketPoolsHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		poolsWithData, err := blockchain.MarketPools(true, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, poolsWithData, nil)
	}
}
