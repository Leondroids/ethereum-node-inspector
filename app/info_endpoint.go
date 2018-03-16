package app

import (
	"encoding/json"
	"net/http"
)

type InfoResponse struct {
	ClientVersion   string `json:"clientVersion"`
	ProtocolVersion int64  `json:"protocolVersion"`
	NetworkID       string `json:"networkId"`
}

func (it *NodeHandler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	info, err := it.processInfoResponse()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(info)
}

func (it *NodeHandler) processInfoResponse() (*InfoResponse, error) {
	clientVersion, err := it.client.Web3.ClientVersion()
	if err != nil {
		return nil, err
	}

	protocolVersion, err := it.client.Eth.ProtocolVersion()
	if err != nil {
		return nil, err
	}

	networkID, err := it.client.Net.NetVersion()
	if err != nil {
		return nil, err
	}

	return &InfoResponse{
		ClientVersion:   clientVersion,
		ProtocolVersion: protocolVersion,
		NetworkID:       networkID,
	}, nil
}
