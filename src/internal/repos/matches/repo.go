package matches

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "net/http/httputil"

    "github.com/derek-test/blitz-backend/src/internal/entity"
)

func (d *httpClient) GetMatchHistory(ctx context.Context, puuid string, endTimestamp int64, matchType string, count int) ([]string, error) {
    client := &http.Client{}

    queryString := "?"
    if endTimestamp != 0 {
        queryString += "endTime=" + fmt.Sprintf("%d", endTimestamp) + "&"
    }

    if count != 0 {
        queryString += "count=" + fmt.Sprintf("%d", count) + "&"
    }
    if matchType != "" {
        queryString += "type=" + matchType + "&"
    }

    req, err := http.NewRequest("GET", fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids"+queryString, puuid), nil)
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

    var history []string
    err = json.NewDecoder(resp.Body).Decode(&history)
    if err != nil {
        return nil, err
    }

    return history, nil
}

// Cacheable
func (d *httpClient) GetMatch(ctx context.Context, matchID string) (*entity.Match, error) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/%s", matchID), nil)
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

    var g *getMatchResp
    err = json.NewDecoder(resp.Body).Decode(&g)
    if err != nil {
        return nil, err
    }

    res := serializeMatchToEntity(g)
    return res, nil
}
