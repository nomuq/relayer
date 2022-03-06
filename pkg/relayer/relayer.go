package relayer

import (
	"context"
	"fmt"

	"github.com/relayer/relayer/pkg/proto"
	"github.com/relayer/relayer/pkg/store"
	"google.golang.org/grpc/metadata"
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

func (s *RelayerServer) GetAuthToken(context.Context, *proto.GetAuthTokenRequest) (*proto.GetAuthTokenResponse, error) {

	// get isAdmin from metadata
	isAdmin := false
	md, ok := metadata.FromIncomingContext(context.Background())
	if ok {
		if md["is-admin"] != nil {
			isAdmin = true
		}
	}

	fmt.Println("isAdmin:", isAdmin)

	return nil, fmt.Errorf("not implemented")
}
