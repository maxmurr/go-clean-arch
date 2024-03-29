package server

import "context"

type Server interface {
	Start() error
	Shutdown(ctx context.Context, errShutdown chan error)
}
