package matches

import (
	"github.com/derek-test/blitz-backend/src/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

type httpClient struct {
	ApiKey string
}

func NewMatchesRepo(config *config.Config) MatchesRepo {
	return &httpClient{
		ApiKey: config.ApiKey,
	}
}
