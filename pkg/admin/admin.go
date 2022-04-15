/*
 * (C) Copyright 2022 Satish Babariya (https://satishbabariya.com/) and others.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Contributors:
 *     satish babariya (satish.babariya@gmail.com)
 *
 */

package admin

import (
	"github.com/relayer/relayer/pkg/config"
	"github.com/relayer/relayer/pkg/proto"
	"github.com/relayer/relayer/pkg/store"
)

type RelayerAdminServer struct {
	proto.UnimplementedRelayerAdminServer
	config *config.RelayerConfig
	store  *store.Store
}

func NewRelayerAdminServer(config *config.RelayerConfig, store *store.Store) *RelayerAdminServer {
	return &RelayerAdminServer{
		config: config,
		store:  store,
	}
}

// func (s *RelayerServer) ValidateAPISecret(ctx context.Context) error {
// 	// get metadata from context
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return fmt.Errorf("metadata is not provided")
// 	}

// 	// validate api-key
// 	if md.Get("api-key")[0] != s.config.APIKey {
// 		return fmt.Errorf("api-key is not valid")
// 	}

// 	// validate api-secret
// 	if md.Get("api-secret") != nil {
// 		if md.Get("api-secret")[0] != s.config.APISecret {
// 			return fmt.Errorf("unauthorized")
// 		}
// 	}

// 	return nil
// }

// func (s *RelayerServer) GetAuthToken(ctx context.Context, req *proto.GetAuthTokenRequest) (*proto.GetAuthTokenResponse, error) {

// 	// validate api-secret
// 	if err := s.ValidateAPISecret(ctx); err != nil {
// 		return nil, err
// 	}

// 	// generate auth token

// 	return nil, fmt.Errorf("not implemented")
// }
