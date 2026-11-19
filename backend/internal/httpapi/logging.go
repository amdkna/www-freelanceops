package httpapi

import (
 "log/slog"
 "net/http"
 "time"
)
type responseWriter struct{http.ResponseWriter; status int}
func(w *responseWriter)WriteHeader(status int){if w.status==0{w.status=status;w.ResponseWriter.WriteHeader(status)}}
func(w *responseWriter)Write(p []byte)(int,error){if w.status==0{w.WriteHeader(200)};return w.ResponseWriter.Write(p)}
func RequestLog(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
 start:=time.Now()
 recorder:=&responseWriter{ResponseWriter:w}
 next.ServeHTTP(recorder,r)
 status:=recorder.status;if status==0{status=200}
 // Only log method and route pattern; never full URLs, query strings, headers, or bodies.
 slog.Info("http request","method",r.Method,"status",status,"duration_ms",time.Since(start).Milliseconds())
})}
