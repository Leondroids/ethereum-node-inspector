package app

import (
	"encoding/json"
	"net/http"
	"github.com/Leondroids/go-ethereum-rpc/types"
)

type AccountsResponse struct {
	Coinbase string `json:"coinbase"`
	Accounts []string `json:"accounts"`
}

func (it *NodeHandler) Accounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	accounts, err := it.processAccounts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(accounts)
}

func (it *NodeHandler) processAccounts() (*AccountsResponse, error) {
	response := &AccountsResponse{}
	var err error

	// accounts
	list, err := it.client.Personal.ListAccounts()
	if err != nil {
		return nil, err
	}

	response.Accounts = types.HexStringListToStringList(list)

	// coinbase
	coinbase, err := it.client.Eth.Coinbase()
	if err != nil {
		return nil, err
	}

	response.Coinbase = coinbase.String()

	return response, nil
}
