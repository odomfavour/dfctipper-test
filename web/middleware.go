package web

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

// CacheControl creates a new middleware to set the HTTP response header with
// "Cache-Control: max-age=maxAge" where maxAge is in seconds.
func CacheControl(maxAge int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(maxAge, 10))
			next.ServeHTTP(w, r)
		})
	}
}

func (s Server) RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := s.GetUserIDTokenCtx(r)
		if uid == 0 {
			if strings.Contains(r.Header.Get("content-type"), "json") {
				RenderErrorfJSON(w, "Id token not available")
			} else {
				NotifyError(w, "Please login to continue")
				http.Redirect(w, r, "/signin?r="+r.RequestURI, http.StatusSeeOther)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ContractAddressCtx returns a http.HandlerFunc that embeds the value at the url
// part {contractAddress} into the request context.
func ContractAddressCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), CtxContractAddress,
			chi.URLParam(r, "contractAddress"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetContractAddressCtx(r *http.Request) string {
	chartAxisType, ok := r.Context().Value(CtxContractAddress).(string)
	if !ok {
		log.Trace("chart axis type not set")
		return ""
	}
	return chartAxisType
}

func (s Server) GetUserIDTokenCtx(r *http.Request) int {
	session, err := s.SessionStore.Get(r, "tth")
	if err != nil {
		return 0
	}
	uid := session.Values["uid"]
	if uid == nil || uid.(int) == 0 {
		return 0
	}

	return uid.(int)
}
