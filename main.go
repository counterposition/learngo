package main

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/samber/do"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const serverAddress = "localhost:8080"
const redisAddress = "localhost:6379"

func main() {
	container := do.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		<-c
		if err := container.Shutdown(); err != nil {
			slog.Error("unable to cleanly shut down the DI container")
		}
	}()

	do.ProvideValue(container, slog.Default())
	do.Provide(container, newHttpServer)
	do.Provide(container, newHttpHandler)
	do.Provide(container, newRedis)
	server := do.MustInvoke[*httpServer](container)

	_ = server.Start()
}

type httpServer struct {
	log    *slog.Logger
	server *http.Server
}

func (hs *httpServer) Start() error {
	hs.log.Info("starting http server", "address", hs.server.Addr)
	return hs.server.ListenAndServe()
}

func (hs *httpServer) Shutdown() error {
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.TODO(), deadline)
	defer cancel()

	if err := hs.server.Shutdown(ctx); err != nil {
		hs.log.Error("server did not shut down normally", "error", err)
		return err
	}

	return nil
}

func newHttpServer(container *do.Injector) (*httpServer, error) {
	log := do.MustInvoke[*slog.Logger](container)
	handler, err := do.Invoke[http.Handler](container)
	if err != nil {
		log.Error("http handler not provided")
		return nil, err
	}
	mux := http.NewServeMux()
	mux.Handle("/", handler)

	server := &http.Server{
		Addr:    serverAddress,
		Handler: mux,
	}

	log.Info("instantiating HTTP server")
	return &httpServer{log: log, server: server}, nil
}

func newRedis(_ *do.Injector) (*miniredis.Miniredis, error) {
	server := miniredis.NewMiniRedis()
	if err := server.StartAddr(redisAddress); err != nil {
		return nil, err
	}

	return server, nil
}

type httpHandler struct {
	log   *slog.Logger
	redis *miniredis.Miniredis
}

func (h *httpHandler) Shutdown() error {
	h.log.Info("shutting down http handler")
	h.redis.Close()
	return nil
}

func (h *httpHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	count, err := h.redis.Incr("count", 1)
	if err != nil {
		h.log.Error("failed to increment count")
		return
	}
	fmt.Fprintf(writer, "received %d requests thus far\n", count)
}

func newHttpHandler(container *do.Injector) (http.Handler, error) {
	log := do.MustInvoke[*slog.Logger](container)
	redis := do.MustInvoke[*miniredis.Miniredis](container)
	return &httpHandler{
		log:   log,
		redis: redis,
	}, nil
}
