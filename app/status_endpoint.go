package app

import (
	"github.com/Leondroids/go-ethereum-rpc/rpc"
	"encoding/json"
	"net/http"
	"log"
)

type StatusResponse struct {
	SyncStatus  *rpc.SyncStatus `json:"syncStatus"`
	PeerCount   int64           `json:"peerCount"`
	IsListening bool            `json:"isListening"`
	GasPrice    string          `json:"gasPrice"`
	IsMining    bool            `json:"isMining"`
	Hashrate    int64           `json:"hashrate"`
}

func (it *NodeHandler) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	status, err := it.processStatusResponse()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}

func (it *NodeHandler) processStatusResponse() (*StatusResponse, error) {
	var err error
	status := &StatusResponse{}
	if status.SyncStatus, err = it.client.Eth.Syncing(); err != nil {
		return nil, err
	}
	if status.IsListening, err = it.client.Net.NetListening(); err != nil {
		return nil, err
	}
	if status.IsMining, err = it.client.Eth.Mining(); err != nil {
		return nil, err
	}
	if status.Hashrate, err = it.client.Eth.Hashrate(); err != nil {
		return nil, err
	}
	if status.PeerCount, err = it.client.Net.NetPeerCount(); err != nil {
		return nil, err
	}
	gp, err := it.client.Eth.GasPrice()
	if err != nil {
		return nil, err
	}
	status.GasPrice = gp.String()

	return status, nil
}
