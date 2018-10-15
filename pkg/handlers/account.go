package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-common/pkg/blockchain"
	"github.com/gorilla/mux"
)

func AccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]
	address := vars["address"]
	var symbolEnum blockchain.BalanceType

	if symbol == "gla" {
		symbolEnum = blockchain.GLA
	} else if symbol == "eth" {
		symbolEnum = blockchain.ETH
	} else {
		symbolNotFoundErr := errors.New("symbol not found for " + symbol)
		ErrorHandler(w, r, "Symbol not supported at this time, try `eth` or `gla`", symbolNotFoundErr, http.StatusNotFound)
		return
	}

	balance, err := blockchain.GetAccountBalance(common.HexToAddress(address), blockchain.BalanceType(symbolEnum))

	if err != nil {
		ErrorHandler(w, r, "Could not retrieve balance for "+address, err, http.StatusInternalServerError)
		return
	}

	ResponseHandler(w, r, "null", true, nil, balance, nil)
}

func AccountTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	decoder := json.NewDecoder(r.Body)
	var options blockchain.TransactionOptions
	err := decoder.Decode(&options)

	transactions, err := blockchain.GetAccountTransactions(common.HexToAddress(address), options)

	if err != nil {
		ErrorHandler(w, r, "Could not retrieve transactions for "+address, err, http.StatusInternalServerError)
		return
	}

	ResponseHandler(w, r, "null", true, nil, transactions.Transactions, nil)
}


func AccountNotFoundErrorHandler(w http.ResponseWriter, r *http.Request, ga *blockchain.GladiusAccountManager) error {
	if !ga.HasAccount() {
		err := errors.New("account not found")
		ErrorHandler(w, r, "Account not found, please create an account", err, http.StatusBadRequest)
		return err
	}

	return nil
}

func AccountUnlockedErrorHandler(w http.ResponseWriter, r *http.Request, ga *blockchain.GladiusAccountManager) error {
	if !ga.Unlocked() {
		err := errors.New("wallet locked")
		ErrorHandler(w, r, "Wallet could not be opened, passphrase is incorrect", err, http.StatusMethodNotAllowed)
		return err
	}
	return nil
}

// Account Manager Error Handler, checks required account permissions prior to accessing API endpoints
func AccountErrorHandler(w http.ResponseWriter, r *http.Request, ga *blockchain.GladiusAccountManager) error {
	err := AccountNotFoundErrorHandler(w, r, ga)
	if err != nil {
		return err
	}

	err = AccountUnlockedErrorHandler(w, r, ga)
	if err != nil {
		return err
	}

	return nil
}

