package relayer

import (
	"github.com/relayer/relayer/pkg/proto"
	"github.com/relayer/relayer/pkg/store"
)

type RelayerServer struct {
	proto.UnimplementedRelayerServer

	store *store.Store
}

func NewRelayerServer(store *store.Store) *RelayerServer {
	return &RelayerServer{
		store: store,
	}
}
