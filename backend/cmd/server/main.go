package main

import(
 "context"
 "errors"
 "log/slog"
 "net"
 "net/http"
 "os"
 "os/signal"
 "syscall"
 "time"

 "github.com/amdkna/www-freelanceops/backend/internal/auth"
 "github.com/amdkna/www-freelanceops/backend/internal/config"
 "github.com/amdkna/www-freelanceops/backend/internal/httpapi"
 "github.com/amdkna/www-freelanceops/backend/internal/storage"
)
func main(){if err:=run();err!=nil{slog.Error("server failed","error_type","startup");os.Exit(1)}}
func run()error{
 cfg,err:=config.Load();if err!=nil{return err}
 level:=map[string]slog.Level{"debug":slog.LevelDebug,"info":slog.LevelInfo,"warn":slog.LevelWarn,"error":slog.LevelError}[cfg.LogLevel]
 slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{Level:level})))
 ctx:=context.Background()
 db,err:=storage.Open(ctx,cfg.DBPath);if err!=nil{return err};defer db.Close()
 authStore:=auth.New(db,cfg.PublicOrigin)
 if err:=authStore.Bootstrap(ctx,cfg.AdminPassword);err!=nil{return err}
 listener,err:=net.Listen("tcp",cfg.Addr);if err!=nil{return err};defer listener.Close()
 health:=httpapi.NewHealth()
 server:=&http.Server{Handler:httpapi.RequestLog(httpapi.RouterWithAuth(health,authStore)),ReadHeaderTimeout:5*time.Second,IdleTimeout:60*time.Second}
 ctx,stop:=signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM);defer stop()
 serverErrors:=make(chan error,1)
 go func(){serverErrors<-server.Serve(listener)}()
 slog.Info("server listening","address",listener.Addr().String())
 select{
 case err:=<-serverErrors:if errors.Is(err,http.ErrServerClosed){return nil};return err
 case <-ctx.Done():
  health.SetReady(false);slog.Info("server shutting down")
  shutdownCtx,cancel:=context.WithTimeout(context.Background(),cfg.ShutdownTimeout);defer cancel()
  if err:=server.Shutdown(shutdownCtx);err!=nil{_ = server.Close();return err}
  if err:=<-serverErrors;err!=nil&&!errors.Is(err,http.ErrServerClosed){return err}
  slog.Info("server stopped");return nil
 }
}
