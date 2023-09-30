package service

import "context"

type Client interface {
	Notify(ctx context.Context, message string) error
}
