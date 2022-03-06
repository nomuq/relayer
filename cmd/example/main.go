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

package main

import (
	"context"

	"github.com/relayer/relayer/pkg/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	interceptor := &ClientInterceptor{
		apiKey:    "cXjnRFK8Naep171",
		apiSecret: "V59vTKSOahyEGgaOzYKYm5m4tnoReDGuxaRMnkMeVm5hSFXCZtYAqQMyt7ZM",
	}

	conn, err := grpc.Dial(
		"localhost:1203",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(interceptor.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(interceptor.StreamClientInterceptor),
	)
	if err != nil {
		logrus.Fatalf("failed to dial: %v", err)
	}
	client := proto.NewRelayerClient(conn)
	resp, err := client.GetAuthToken(
		context.Background(),
		&proto.GetAuthTokenRequest{},
	)
	if err != nil {
		logrus.Fatalf("failed to call: %v", err)
	}
	logrus.Infof("resp: %v", resp)
}

// ClientInterceptor is a gRPC interceptor that adds the access token to the request
type ClientInterceptor struct {
	apiKey      string
	apiSecret   string
	accessToken string
}

// UnaryClientInterceptor is a gRPC interceptor that adds the access token to the request
func (interceptor *ClientInterceptor) UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	if interceptor.apiKey != "" {
		md.Set("api-key", interceptor.apiKey)
	}

	if interceptor.apiSecret != "" {
		md.Set("api-secret", interceptor.apiSecret)
	}

	if interceptor.accessToken != "" {
		md.Set("authorization", interceptor.accessToken)
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}

// StreamClientInterceptor is a gRPC interceptor that adds the access token to the request
func (interceptor ClientInterceptor) StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	if interceptor.accessToken != "" {
		md.Set("authorization", interceptor.accessToken)
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return streamer(ctx, desc, cc, method, opts...)
}
