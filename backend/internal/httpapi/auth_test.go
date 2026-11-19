package httpapi

import(
 "bytes"
 "encoding/json"
 "net/http"
 "net/http/httptest"
 "path/filepath"
 "strings"
 "testing"

 "github.com/amdkna/www-freelanceops/backend/internal/auth"
 "github.com/amdkna/www-freelanceops/backend/internal/storage"
 "context"
)
func TestAuthenticationCSRFAndLogout(t *testing.T){
 db,err:=storage.Open(context.Background(),filepath.Join(t.TempDir(),"auth.db"));if err!=nil{t.Fatal(err)};defer db.Close()
 s:=auth.New(db,"https://example.com")
 if err:=s.Bootstrap(context.Background(),"LongExamplePassword123!");err!=nil{t.Fatal(err)}
 h:=RouterWithAuth(NewHealth(),s)
 send:=func(method,path,body,origin,csrf string,cookie *http.Cookie)*httptest.ResponseRecorder{
  t.Helper();req:=httptest.NewRequest(method,path,strings.NewReader(body))
  if origin!=""{req.Header.Set("Origin",origin)}
  if csrf!=""{req.Header.Set("X-CSRF-Token",csrf)}
  if cookie!=nil{req.AddCookie(cookie)}
  w:=httptest.NewRecorder();h.ServeHTTP(w,req);return w
 }
 if w:=send("POST","/v1/auth/login",`{"password":"bad"}`,"https://example.com","",nil);w.Code!=401{t.Fatalf("wrong password=%d",w.Code)}
 if w:=send("POST","/v1/auth/login",`{"password":"LongExamplePassword123!"}`,"https://evil.example","",nil);w.Code!=403{t.Fatalf("wrong origin=%d",w.Code)}
 w:=send("POST","/v1/auth/login",`{"password":"LongExamplePassword123!"}`,"https://example.com","",nil)
 if w.Code!=200{t.Fatalf("login=%d %s",w.Code,w.Body.String())}
 cookies:=w.Result().Cookies();if len(cookies)==0{t.Fatal("session cookie missing")}
 cookie:=cookies[0]
 if !cookie.Secure||!cookie.HttpOnly||cookie.SameSite!=http.SameSiteStrictMode{t.Fatal("unsafe cookie")}
 var resp struct{CSRF string `json:"csrf_token"`}
 if err:=json.NewDecoder(bytes.NewReader(w.Body.Bytes())).Decode(&resp);err!=nil{t.Fatal(err)}
 if w:=send("GET","/v1/auth/session","","","",cookie);w.Code!=200{t.Fatalf("session=%d",w.Code)}
 if w:=send("GET","/v1/websites","","","",nil);w.Code!=401{t.Fatalf("unauthenticated=%d",w.Code)}
 if w:=send("POST","/v1/websites","{}","https://example.com","",cookie);w.Code!=403{t.Fatalf("missing csrf=%d",w.Code)}
 if w:=send("POST","/v1/websites","{}","https://evil.example",resp.CSRF,cookie);w.Code!=403{t.Fatalf("wrong origin=%d",w.Code)}
 if w:=send("POST","/v1/websites","{}","https://example.com",resp.CSRF,cookie);w.Code!=501{t.Fatalf("protected placeholder=%d",w.Code)}
 if w:=send("POST","/v1/auth/logout","","https://example.com",resp.CSRF,cookie);w.Code!=204{t.Fatalf("logout=%d",w.Code)}
 if w:=send("GET","/v1/auth/session","","","",cookie);w.Code!=401{t.Fatalf("revoked=%d",w.Code)}
}
