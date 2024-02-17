package matches

import (
	"github.com/derek-test/blitz-backend/src/internal/cache"
	"github.com/derek-test/blitz-backend/src/internal/repos/matches"
	"github.com/derek-test/blitz-backend/src/internal/repos/summoner"
)

type service struct {
	matchRepo    matches.MatchesRepo
	summonerRepo summoner.SummonerRepo
	cache        cache.Cache
}

func NewMatchService(matchRepo matches.MatchesRepo, summonerRepo summoner.SummonerRepo, cache cache.Cache) MatchService {
	return &service{
		summonerRepo: summonerRepo,
		matchRepo:    matchRepo,
		cache:        cache,
	}
}
