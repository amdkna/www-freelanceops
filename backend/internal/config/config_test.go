package config

import("testing";"time")
func TestLoad(t *testing.T){
 for _,tc:=range []struct{name,port,origin,db,timeout,level string; bad bool}{
  {"missing port","","https://example.com","","","",true},
  {"invalid port","abc","https://example.com","","","",true},
  {"invalid timeout","8080","https://example.com","","0s","",true},
  {"missing origin","8080","","","","",true},
  {"bad origin","8080","file:///tmp","","","",true},
  {"relative db","8080","https://example.com","data/db","","",true},
  {"invalid level","8080","https://example.com","","","verbose",true},
  {"valid","8080","https://example.com","/tmp/test.db","2s","info",false},
 }{
  t.Run(tc.name,func(t *testing.T){
   t.Setenv("PORT",tc.port);t.Setenv("PUBLIC_ORIGIN",tc.origin);t.Setenv("DB_PATH",tc.db);t.Setenv("SHUTDOWN_TIMEOUT",tc.timeout);t.Setenv("LOG_LEVEL",tc.level)
   got,err:=Load();if (err!=nil)!=tc.bad{t.Fatalf("Load err=%v expected bad=%v",err,tc.bad)}
   if !tc.bad&&(got.Addr!=":8080"||got.ShutdownTimeout!=2*time.Second){t.Fatalf("unexpected config: %+v",got)}
  })
 }
}
