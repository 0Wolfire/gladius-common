package routing

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/gladiusio/gladius-common/pkg/blockchain"
	"github.com/gladiusio/gladius-common/pkg/handlers"
	"github.com/gladiusio/gladius-p2p/pkg/p2p/peer"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ControlRouter struct {
	Router *mux.Router
	Port   string
	Debug  bool
}

func (cRouter *ControlRouter) Start() {
	if cRouter.Debug {
		cRouter.Router.Use(loggingMiddleware)
	}

	fmt.Println("Starting API at http://localhost:" + cRouter.Port)
	log.Fatal(http.ListenAndServe(":"+cRouter.Port, ghandlers.CORS()(cRouter.Router)))
}

func AppendP2PEndPoints(router *mux.Router, ga *blockchain.GladiusAccountManager) error {
	// P2P setup
	peerStruct := peer.New(ga)
	p2pRouter := router.PathPrefix("/p2p").Subrouter()
	// P2P Message Routes
	p2pRouter.HandleFunc("/message/sign", handlers.CreateSignedMessageHandler(ga)).
		Methods(http.MethodPost)
	p2pRouter.HandleFunc("/message/verify", handlers.VerifySignedMessageHandler).
		Methods("POST")

	p2pRouter.HandleFunc("/network/join", handlers.JoinHandler(peerStruct)).
		Methods("POST")

	p2pRouter.HandleFunc("/network/leave", handlers.LeaveHandler(peerStruct)).
		Methods("POST")

	// P2P State Routes
	p2pRouter.HandleFunc("/state/push_message", handlers.PushStateMessageHandler(peerStruct)).
		Methods("POST")
	p2pRouter.HandleFunc("/state", handlers.GetFullStateHandler(peerStruct)).
		Methods("GET")
	p2pRouter.HandleFunc("/state/{node_address}", handlers.GetNodeStateHandler(peerStruct)).
		Methods("GET")
	p2pRouter.HandleFunc("/state/signatures", handlers.GetSignatureListHandler(peerStruct)).
		Methods("GET")
	p2pRouter.HandleFunc("/state/content_diff", handlers.GetContentNeededHandler(peerStruct)).
		Methods("POST")
	p2pRouter.HandleFunc("/state/content_links", handlers.GetContentLinksHandler(peerStruct)).
		Methods("POST")

	// Only enable for testing
	if viper.GetBool("NodeManager.Config.Debug") {
		p2pRouter.HandleFunc("/state/set_state", handlers.SetStateDebugHandler(peerStruct)).
			Methods("POST")
	}

	return nil
}

func AppendAccountManagementEndpoints(router *mux.Router) error {
	// Account Management
	accountRouter := router.PathPrefix("/account/{address:0[xX][0-9a-fA-F]{40}}").Subrouter().StrictSlash(true)
	accountRouter.HandleFunc("/balance/{symbol:[a-z]{3}}", handlers.AccountBalanceHandler)
	accountRouter.HandleFunc("/transactions/{symbol:[a-z]{3}}", handlers.AccountTransactionsHandler)

	return nil
}

func AppendWalletManagementEndpoints(router *mux.Router, ga *blockchain.GladiusAccountManager) error {
	// Key Management
	walletRouter := router.PathPrefix("/keystore").Subrouter().StrictSlash(true)
	walletRouter.HandleFunc("/account/create", handlers.KeystoreAccountCreationHandler(ga)).
		Methods(http.MethodPost)
	walletRouter.HandleFunc("/account", handlers.KeystoreAccountRetrievalHandler(ga))
	walletRouter.HandleFunc("/account/open", handlers.KeystoreAccountUnlockHandler(ga)).
		Methods(http.MethodPost)

	return nil
}

func AppendStatusEndpoints(router *mux.Router) error {
	// TxHash Status Sub-Routes
	statusRouter := router.PathPrefix("/status").Subrouter().StrictSlash(true)
	statusRouter.HandleFunc("/", handlers.StatusHandler).
		Methods(http.MethodGet, http.MethodPut).
		Name("status")
	statusRouter.HandleFunc("/tx/{tx:0[xX][0-9a-fA-F]{64}}", handlers.StatusTxHandler).
		Methods(http.MethodGet).
		Name("status-tx")

	return nil
}

func AppendVersionEndpoints(router *mux.Router) error {
	versionRouter := router.PathPrefix("/version").Subrouter()
	versionRouter.HandleFunc("/version", handlers.VersionHandler()).Methods("GET")

	return nil
}

func responseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println()
		log.Println(formatRequest(r))
		log.Println()

		next.ServeHTTP(w, r)
	})
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string

	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == http.MethodPost {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	// Return the request as a string
	return strings.Join(request, "\n")
}
