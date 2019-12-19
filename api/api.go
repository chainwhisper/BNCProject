package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/binance-chain/go-sdk/client/rpc"
	ctypes "github.com/binance-chain/go-sdk/common/types"
	newtypes "github.com/binance-chain/go-sdk/types"
	"github.com/binance-chain/go-sdk/types/tx"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// Server represents the API server
type Server struct {
	Port int `json:"port"`

	Url string
	Version string
	Commit  string
	Branch  string
}
type encodeResp struct {
	Tx string `json:"tx"`
}

// Router returns the router
func (s *Server) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/version", s.VersionHandler)
	router.HandleFunc("/tx", s.TxHashHandler).Methods("POST")
	return router
}


// VersionHandler handles the /version route
func (s *Server) VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("{\"version\": \"%s\"}", s.Version)))
}

// TxHandler handles the /tx route
func (s *Server) TxHashHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("get request",string(body), " ",len(body) )

	testClientInstance := rpc.NewRPCClient(s.Url, ctypes.TestNetwork)

	//bz, err := hex.DecodeString(testTxHash)
	input := string(body)
	bz, err := hex.DecodeString(input)

	txResult, err := testClientInstance.Tx(bz, false)

	if err != nil{

		fmt.Println("error: ",err)
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	codec := newtypes.NewCodec()
	parsedTx :=  tx.StdTx{}

	err = codec.UnmarshalBinaryLengthPrefixed(txResult.Tx, &parsedTx)

	bz, err = codec.MarshalJSON(parsedTx)
	fmt.Println(string(bz))

	// Return to client
	output, err := json.Marshal(encodeResp{Tx: string(bz)})
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}

func writeErrorResponse(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	bz,_ :=json.Marshal(errorResponse{Code: 0, Message: err})
	w.Write(bz)
}

// errorResponse defines the attributes of a JSON error response.
type errorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}