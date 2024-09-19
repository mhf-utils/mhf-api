package middlewares

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	_ "github.com/lib/pq"

	"github.com/newrelic/go-agent/v3/newrelic"

	"mhf-api/utils/logger"
)

var context_middlewares_logger = "api:middlewares:logger"

type customResponseWriter struct {
	http.ResponseWriter
	buf bytes.Buffer
}

func Logging(log *logger.Logger, newRelicApp *newrelic.Application) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			logWithContext := log.WithContext(context_middlewares_logger, "method:Logging")
			clientIP := req.RemoteAddr
			endpoint := req.URL.Path
			userAgent := req.UserAgent()
			txn := newRelicApp.StartTransaction(endpoint)

			defer txn.End()

			logWithContext.Info(fmt.Sprintf("Request received - Endpoint: %s | Method: %s | ClientIP: %s | UserAgent: %s", endpoint, req.Method, clientIP, userAgent))

			customRes := &customResponseWriter{ResponseWriter: res, buf: bytes.Buffer{}}

			defer func() {
				if strings.Contains(endpoint, "/launcher/files") {
					logWithContext.Info(fmt.Sprintf("Response sent - Endpoint: %s", endpoint))
				} else {
					logWithContext.Info(fmt.Sprintf("Response sent - Endpoint: %s | Body: %s", endpoint, customRes.buf.String()))
				}
			}()

			body, err := io.ReadAll(req.Body)
			if err != nil {
				logWithContext.Error(fmt.Sprintf("Error reading body: %s", err.Error()))
			}

			req.Body = io.NopCloser(bytes.NewBuffer(body))

			logWithContext.Info(fmt.Sprintf("URI='%s' METHOD=['%s'] QUERY=%s BODY=%s", req.RequestURI, req.Method, req.URL.Query(), body))

			txn.SetWebRequestHTTP(req)
			txn.SetWebResponse(customRes)
			next.ServeHTTP(customRes, req)
		})
	}
}

func (custom *customResponseWriter) Write(b []byte) (int, error) {
	n, err := custom.buf.Write(b)
	if err != nil {
		return n, errors.New("failed to write response data to buffer")
	}
	_, err = custom.ResponseWriter.Write(b)
	if err != nil {
		return n, errors.New("failed to write response data to original ResponseWriter")
	}
	return n, nil
}

func (custom *customResponseWriter) WriteHeader(statusCode int) {
	custom.ResponseWriter.WriteHeader(statusCode)
}
