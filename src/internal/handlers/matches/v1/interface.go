package v1

import "net/http"

type MatchesHandler interface {
    GetMatchHistory(w http.ResponseWriter, r *http.Request)
}
