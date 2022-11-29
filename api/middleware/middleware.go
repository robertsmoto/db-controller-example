package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/robertsmoto/db_controller_example/config"
	"github.com/robertsmoto/db_controller_example/repo/models"
	"github.com/robertsmoto/db_controller_example/repo/redisdb"
)

func TimerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// sets the start time of the request
		t1 := time.Now()
		ctx := context.WithValue(r.Context(), "startTime", t1)
		println("## middle 01")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acct := new(models.Account)
		acct.ID = r.Header.Get("Aid")
		acct.Auth = r.Header.Get("Auth")
		acct.Prefix = r.Header.Get("Prefix")
		if acct.ID == "" || acct.Auth == "" || acct.Prefix == "" {
			msg := strings.NewReader("")
			msgReadCloser := io.NopCloser(msg)
			r.Body = msgReadCloser
			http.Error(w, `{"errors":"not authorized"}`, http.StatusForbidden)
		}
		ctx1 := context.WithValue(r.Context(), "Acct", acct)
		ctx2 := context.WithValue(ctx1, "Conf", config.Conf)
		println("## middle 02")
		next.ServeHTTP(w, r.WithContext(ctx2))
	})
}

func (h *MiddlewareConn) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// information passed from client in the request header
		var err error
		ctx := r.Context()
		conf := ctx.Value("Conf").(*config.Config)
		acct := ctx.Value("Acct").(*models.Account)
		hitsKey := fmt.Sprintf("%s:hits", acct.ID)
		holdKey := fmt.Sprintf("%s:hold", acct.ID)
		holdCntKey := fmt.Sprintf("%s:cntr", acct.ID)
		rds := redisdb.ConnectDB(redisdb.Api)
		defer rds.Close()
		apiHits, err := rds.Do("INCR", hitsKey)
		if apiHits.(int64) == 1 {
			_, err = rds.Do("EXPIRE", hitsKey, conf.ThresholdTime)
		}
		if apiHits.(int64) >= conf.ThresholdHits {
			// reset the hits key
			_, err = rds.Do("SET", hitsKey, 1)
			_, err = rds.Do("EXPIRE", hitsKey, conf.ThresholdTime)
			// incr hold counter
			holdCntr, _ := rds.Do("INCR", holdCntKey)
			if holdCntr.(int64) == 1 {
				_, err = rds.Do("EXPIRE", holdCntKey, (60 * 60 * 24))
			}
			// this sets and increments a holdKey
			_, err = rds.Do("INCR", holdKey)
			// ttl for holdKey based on exponent of hold counter
			holdTime := int(math.Pow(
				float64(conf.ThresholdTime), float64(holdCntr.(int64))))
			_, err = rds.Do("EXPIRE", holdKey, holdTime)
			if err != nil {
				log.Fatal(err)
			}
		}
		// check if holdKey exists
		holdExists, _ := rds.Do("GET", holdKey)
		defer rds.Close()
		if holdExists != nil {
			msg := strings.NewReader("")
			msgReadCloser := io.NopCloser(msg)
			r.Body = msgReadCloser
			http.Error(w, `{"errors"
      :"api limit reached, access denied"}`, http.StatusForbidden)
		} else {
			println("## middle 03")
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

// authorize account user
func (h *MiddlewareConn) AccountAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// information passed from client in the request header
		ctx := r.Context()
		// get request account
		rqstAcct := ctx.Value("Acct").(*models.Account)
		// get repo account
		rds := redisdb.ConnectDB(redisdb.Account)
		result, _ := rds.Do("JSON.GET", rqstAcct.ID)
		// expect result to be []byte
		if result == nil {
			result = []byte("{}")
		}
		repoAcct := new(models.Account)
		json.Unmarshal(result.([]byte), repoAcct)
		//validate the account
		if rqstAcct.Auth != repoAcct.Auth || rqstAcct.Prefix != repoAcct.Prefix {
			msg := strings.NewReader("")
			msgReadCloser := io.NopCloser(msg)
			r.Body = msgReadCloser
			http.Error(w, `{"errors":"not authorized"}`, http.StatusForbidden)
		} else {
			println("## middle 04")
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

// validate the contentType
func ContentAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hCType := r.Header.Get("Content-Type")
		if r.Method == "POST" && hCType != "application/json" {
			msg := strings.NewReader("")
			msgReadCloser := io.NopCloser(msg)
			r.Body = msgReadCloser
			http.Error(w, `{"errors":"content type not allowed"}`, http.StatusForbidden)
		}
		println("## middle 05")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
