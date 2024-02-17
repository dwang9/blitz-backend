package v1

import (
	"github.com/derek-test/blitz-backend/src/internal/service/matches"
)

type handler struct {
	matchesService matches.MatchService
}

func NewMatchHandler(matchesService matches.MatchService) MatchesHandler {
	return &handler{
		matchesService: matchesService,
	}
}
