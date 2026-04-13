package hello

import "context"

type Service struct{}

func (s *Service) SayHello(ctx context.Context, name string) (string, error) {
	return "Hello, " + name, nil
}