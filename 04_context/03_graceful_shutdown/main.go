package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type crawlerImpl struct {
	client  *http.Client
	cacheMx *sync.RWMutex
	cache   map[string]cacheEntry
}

const (
	shutdownTimeout = 10 * time.Second
	workersLimit    = 1 << 10
	cacheTTL        = time.Second
)

func New() *crawlerImpl {
	return &crawlerImpl{
		cache:   make(map[string]cacheEntry),
		cacheMx: new(sync.RWMutex),
		client: &http.Client{
			//nolint:mnd // http.Transport defaults tuning
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				MaxIdleConnsPerHost:   runtime.GOMAXPROCS(-1),
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: time.Second * 3,
				ResponseHeaderTimeout: time.Minute,
			},
		},
	}
}

func (c *crawlerImpl) ListenAndServe(ctx context.Context, address string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /crawl", c.crawlHandler)

	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		// or time.Sleep
		if err := srv.Shutdown(shutdownCtx); err != nil {
			// We can join errors
			log.Printf("http server shutdown error: %v", err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (c *crawlerImpl) crawlHandler(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

type cacheEntry struct {
	expiredAt time.Time
}
