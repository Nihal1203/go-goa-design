package user

import "context"

type Service struct {
}

func (u *Service) GetUser(ctx context.Context, id string) (string, error) {
	return id + "1", nil
}
