package repo

import (
	"fmt"
	"muzz_challenge/cmd/internal/types"
)

type Repo struct {
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) ListLikedYou(userId string, pageSize int, paginationToken *string) ([]*types.User, *string, error) {
	return nil, nil, fmt.Errorf("not implemented")
}

func (r *Repo) PutDecision(deciderId string, decision *types.Decision) error {
	return fmt.Errorf("not implemented")
}

func (r *Repo) ListNewLikedYou(
	userId string,
	paginationToken *string,
) ([]*types.User, *string, error) {
	return nil, nil, fmt.Errorf("not implemented")
}

func (r *Repo) CountLikedYou(userId string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
