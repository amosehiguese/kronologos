package server

import (
	api "github.com/amosehiguese/proglog/api/v1"
)

type Config struct {
	CommitLog CommitLog
}

var _ api.LogServer = (*grpcServer)(nil)

type grpcSever struct {
	api.UnimplementedLogServer
	*Config
}

func newgrpcServer(config *Config) (srv *grpcSever, err error) {
	srv = &grpcSever{
		Config: config,
	}
	return srv, nil
}