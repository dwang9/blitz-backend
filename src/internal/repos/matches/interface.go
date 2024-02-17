package matches

import (
	"context"

	"github.com/derek-test/blitz-backend/src/internal/entity"
)

//go:generate mockgen -source interface.go -destination mocks/mock.go -mock_names MatchesRepo=MockMatchesRepo
type MatchesRepo interface {
	GetMatchHistory(ctx context.Context, puuid string, endTimestamp int64, matchType string, count int) ([]string, error)
	GetMatch(ctx context.Context, matchID string) (*entity.Match, error)
}
