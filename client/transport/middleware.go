package transport

/*
   AIMS (Attacked Infrastructure Modular Specification)
   Copyright (C) 2021 Maxime Landon

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/reeflective/team/client"
	"github.com/reeflective/team/example/transports/grpc/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	kb = 1024
	mb = kb * 1024
	gb = mb * 1024

	// ClientMaxReceiveMessageSize - Max gRPC message size ~2Gb.
	ClientMaxReceiveMessageSize = (2 * gb) - 1 // 2Gb - 1 byte

	defaultTimeout = 10 * time.Second
)

// ErrNoTLSCredentials is an error raised if the teamclient was asked to setup, or try
// connecting with, TLS credentials. If such an error is raised, make sure your team
// client has correctly fetched -using client.Config()- a remote teamserver config.
var ErrNoTLSCredentials = errors.New("the Teamclient has no TLS credentials to use")

// TokenAuth extracts authentication metadata from contexts,
// specifically the "Authorization": "Bearer" key:value pair.
type TokenAuth string

// LogMiddlewareOptions is an example list of gRPC options with logging middleware set up.
// This function uses the core teamclient loggers to log the gRPC stack/requests events.
// The Teamclient of this package uses them by default.
func LogMiddlewareOptions(cli *client.Client) []grpc.DialOption {
	logrusEntry := cli.NamedLogger("transport", "grpc")
	logrusOpts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(common.CodeToLevel),
	}

	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	// Intercepting client requests.
	requestIntercept := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		rawRequest, err := json.Marshal(req)
		if err != nil {
			logrusEntry.Errorf("Failed to serialize: %s", err)
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		logrusEntry.Debugf("Raw request: %s", string(rawRequest))

		return invoker(ctx, method, req, reply, cc, opts...)
	}

	options := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(grpc_logrus.UnaryClientInterceptor(logrusEntry, logrusOpts...)),
		grpc.WithUnaryInterceptor(requestIntercept),
	}

	return options
}

func TLSAuthMiddleware(cli *client.Client) ([]grpc.DialOption, error) {
	config := cli.Config()
	if config.PrivateKey == "" {
		return nil, ErrNoTLSCredentials
	}

	tlsConfig, err := cli.NewTLSConfigFrom(config.CACertificate, config.Certificate, config.PrivateKey)
	if err != nil {
		return nil, err
	}

	transportCreds := credentials.NewTLS(tlsConfig)
	callCreds := credentials.PerRPCCredentials(TokenAuth(config.Token))

	return []grpc.DialOption{
		grpc.WithTransportCredentials(transportCreds),
		grpc.WithPerRPCCredentials(callCreds),
	}, nil
}

// Return value is mapped to request headers.
func (t TokenAuth) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + string(t),
	}, nil
}

// RequireTransportSecurity always return true.
func (TokenAuth) RequireTransportSecurity() bool {
	return true
}
