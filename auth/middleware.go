package auth

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//var (
//	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
//		Name: "myapp_processed_ops_total",
//		Help: "The total number of processed events",
//	})
//)

type MetricsMiddleware struct {
	opsProcessed *prometheus.CounterVec
}
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func NewMetricsMiddleware() *MetricsMiddleware {
	opsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"method", "path", "statuscode"})
	return &MetricsMiddleware{
		opsProcessed: opsProcessed,
	}
}
func (lm *MetricsMiddleware) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		wi := &responseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}
		next.ServeHTTP(wi, r)

		lm.opsProcessed.With(prometheus.Labels{"method": r.Method, "path": r.RequestURI, "statuscode": strconv.Itoa(wi.statusCode)}).Inc()
	})
}

//func RecordMetrics(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		wi := &responseWriterInterceptor{w, http.StatusOK}
//		next.ServeHTTP(wi, r)
//	})
//}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorisationHeader := r.Header.Get("authorization")
		if authorisationHeader != "" {
			token, err := jwt.Parse(authorisationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				err := json.NewEncoder(w).Encode(Exception{Message: error.Error(err)})
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
			if token.Valid {
				context.Set(r, "decoded", token.Claims)
				next(w, r)
			} else {
				err := json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		} else {
			err := json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(p []byte) (int, error) {
	return w.ResponseWriter.Write(p)
}

func (w *responseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("type assertion failed http.ResponseWriter not a http.Hijacker")
	}
	return h.Hijack()
}

func (w *responseWriterInterceptor) Flush() {
	f, ok := w.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}
	f.Flush()
}
