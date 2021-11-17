package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gometr/pkg/graceful"
	"net/http"
)

type Handler struct {
	log *zap.Logger
}

func NewHandler(log *zap.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) StartHandle(httpHost, httpPort string) {
	server := &http.Server{
		Addr:    httpHost + ":" + httpPort,
		Handler: h.newRouter(),
	}

	h.log.Info(fmt.Sprintf("Server is running %v", server.Addr))

	graceful.AddCallback(func() error {
		return server.Shutdown(context.Background())
	})

	err := server.ListenAndServe()
	if err != nil {
		h.log.Error("Server shutdown failed", zap.Error(err))
		graceful.ShutdownNow()

		return
	}
}

func (h *Handler) writeResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	payload, err := json.Marshal(response)
	if err != nil {
		h.log.Error("Failed to marshal HTTP response", zap.Error(err), zap.Any("response", response))

		return
	}
	_, err = w.Write(payload)
	if err != nil {
		h.log.Error("Failed to write HTTP response", zap.Error(err), zap.Any("response", response))

		return
	}
}
