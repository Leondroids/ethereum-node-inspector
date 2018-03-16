package app

import "github.com/Leondroids/go-ethereum-rpc/rpc"

type NodeHandler struct {
	client *rpc.Client
}

func NewNodeHandler(appContext *Context) *NodeHandler {
	return &NodeHandler{
		client: appContext.Client,
	}
}

type ResponseString struct {
	Result string `json:"result"`
}
