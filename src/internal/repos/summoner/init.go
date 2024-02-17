package summoner

import (
	"github.com/derek-test/blitz-backend/src/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

type httpClient struct {
	ApiKey string
}

func NewSummonerRepo(config *config.Config) SummonerRepo {
	return &httpClient{
		ApiKey: config.ApiKey,
	}
}
