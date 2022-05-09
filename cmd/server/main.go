package main

import (
	"net/http"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func handler(w http.ResponseWriter, r *http.Request) {
	logger.Info("request received", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message":"hello world!"}`))
}

func main() {
	logger, _ = zap.NewProduction()

	http.HandleFunc("/", handler)

	logger.Info("listening on port :9999")
	logger.Fatal("", zap.Error(http.ListenAndServe(":9999", nil)))
}
