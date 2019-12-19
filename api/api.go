package api


import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/gorilla/mux"
)

type encodeReq struct {
	Tx auth.StdTx `json:"tx"`
}

type encodeResp struct {
	Tx string `json:"tx"`
}

// Server represents the API server
type Server struct {
	Port int `json:"port"`

	Version string
	Commit  string
	Branch  string
}

// Router returns the router
func (s *Server) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/version", s.VersionHandler)

	return router
}


// VersionHandler handles the /version route
func (s *Server) VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("{\"version\": \"%s\", \"commit\": \"%s\", \"branch\": \"%s\"}", s.Version, s.Commit, s.Branch)))
}


// errorResponse defines the attributes of a JSON error response.
type errorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}
