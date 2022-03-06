package relayer

import (
	"context"
	"fmt"

	"github.com/relayer/relayer/pkg/config"
	"github.com/relayer/relayer/pkg/proto"
	"github.com/relayer/relayer/pkg/store"
	"google.golang.org/grpc/metadata"
)

type RelayerServer struct {
	proto.UnimplementedRelayerServer
	config *config.RelayerConfig
	store  *store.Store
}

func NewRelayerServer(config *config.RelayerConfig, store *store.Store) *RelayerServer {
	return &RelayerServer{
		config: config,
		store:  store,
	}
}

func (s *RelayerServer) ValidateAPISecret(ctx context.Context) error {
	// get metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("metadata is not provided")
	}

	// validate api-key
	if md.Get("api-key")[0] != s.config.APIKey {
		return fmt.Errorf("api-key is not valid")
	}

	// validate api-secret
	if md.Get("api-secret") != nil {
		if md.Get("api-secret")[0] != s.config.APISecret {
			return fmt.Errorf("unauthorized")
		}
	}

	return nil
}

func (s *RelayerServer) GetAuthToken(ctx context.Context, req *proto.GetAuthTokenRequest) (*proto.GetAuthTokenResponse, error) {

	// validate api-secret
	if err := s.ValidateAPISecret(ctx); err != nil {
		return nil, err
	}

	// generate auth token

	return nil, fmt.Errorf("not implemented")
}
