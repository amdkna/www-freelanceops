package httpapi

import (
 "encoding/json"
 "net/http"
 "sync/atomic"

 "github.com/amdkna/www-freelanceops/backend/internal/auth"
 "github.com/go-chi/chi/v5"
)
type Health struct{ready atomic.Bool}
func NewHealth()*Health{h:=&Health{};h.ready.Store(true);return h}
func(h *Health)SetReady(ready bool){h.ready.Store(ready)}
func Router(h *Health)http.Handler{return RouterWithAuth(h,nil)}
func RouterWithAuth(h *Health,s *auth.Store)http.Handler{
 r:=chi.NewRouter()
 r.Get("/health/live",func(w http.ResponseWriter,_ *http.Request){writeStatus(w,200,"ok")})
 r.Get("/health/ready",func(w http.ResponseWriter,_ *http.Request){
  if !h.ready.Load(){writeStatus(w,503,"not_ready");return}
  writeStatus(w,200,"ok")
 })
 if s!=nil{authRoutes(r,s)}
 return r
}
func writeStatus(w http.ResponseWriter,code int,status string){
 w.Header().Set("Content-Type","application/json; charset=utf-8");w.Header().Set("Cache-Control","no-store");w.WriteHeader(code)
 _=json.NewEncoder(w).Encode(struct{Status string `json:"status"`}{Status:status})
}
