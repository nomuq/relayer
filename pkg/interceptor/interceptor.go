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

package interceptor

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"

	"github.com/relayer/relayer/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	config *config.RelayerConfig
}

func NewInterceptor(config *config.RelayerConfig) *Interceptor {
	return &Interceptor{
		config: config,
	}
}

func (interceptor *Interceptor) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	err := interceptor.authorize(ctx, info.FullMethod, req)
	if err != nil {
		logrus.Errorf(info.FullMethod, err)
		return nil, err
	}
	logrus.Debugf(info.FullMethod)
	return handler(ctx, req)
}

func (interceptor *Interceptor) StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logrus.Debugf(info.FullMethod)
	err := interceptor.authorize(stream.Context(), info.FullMethod, srv)
	if err != nil {
		logrus.Errorf(info.FullMethod, err)
		return err
	}
	logrus.Debugf(info.FullMethod)
	return handler(srv, stream)
}

func (interceptor *Interceptor) authorize(ctx context.Context, method string, payload interface{}) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	if md.Get("api-key") == nil {
		return status.Errorf(codes.Unauthenticated, "api-key is not provided")
	}

	// validate api-key
	if md.Get("api-key")[0] != interceptor.config.APIKey {
		return status.Errorf(codes.Unauthenticated, "api-key is not valid")
	}

	// if api-secret is provided, validate it else check for authorization header and validate it else return error
	if md.Get("api-secret") != nil {
		if md.Get("api-secret")[0] != interceptor.config.APISecret {
			return status.Errorf(codes.Unauthenticated, "unauthorized")
		}
	} else {
		values := md["authorization"]
		if len(values) == 0 {
			return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		accessToken := values[0]

		token, err := jwt.Parse(accessToken, interceptor.keyFunc)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("claims: ", claims)
			// set claims in metadata
			// md.Set("claims", claims)
		} else {
			return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
		}
	}

	return nil
}

func (interceptor *Interceptor) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(interceptor.config.APISecret), nil
}
