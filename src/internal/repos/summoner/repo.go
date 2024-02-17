package summoner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/derek-test/blitz-backend/src/internal/entity"
)

// Cacheable
func (d *httpClient) GetSummonerInfo(ctx context.Context, summonerName string) (*entity.Summoner, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", summonerName), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Riot-Token", d.ApiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		respErr := errors.New(string(b))
		return nil, respErr
	}

	var s *getSummonerResp
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	res := serializeSummonerToEntity(s)
	return res, nil
}
func (d *httpClient) GetPlayerData(ctx context.Context, puuid string) (*entity.PlayerData, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://na1.api.riotgames.com/lol/challenges/v1/player-data/%s", puuid), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Riot-Token", d.ApiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		respErr := errors.New(string(b))
		return nil, respErr
	}

	var s *playerDataResponse
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	res := serializePlayerDataToEntity(s)
	return res, nil
}
