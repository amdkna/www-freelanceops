package storage

import (
 "context"
 "database/sql"
 "path/filepath"
 "testing"
)
func TestMigrationsPersistenceAndConstraints(t *testing.T){
 ctx:=context.Background();file:=filepath.Join(t.TempDir(),"db.sqlite")
 db,err:=Open(ctx,file);if err!=nil{t.Fatal(err)}
 var mode string
 if err:=db.QueryRow("PRAGMA journal_mode").Scan(&mode);err!=nil||mode!="wal"{t.Fatalf("WAL mode=%q err=%v",mode,err)}
 var foreignKeys int
 if err:=db.QueryRow("PRAGMA foreign_keys").Scan(&foreignKeys);err!=nil||foreignKeys!=1{t.Fatalf("foreign keys=%d err=%v",foreignKeys,err)}
 if _,err:=db.Exec("INSERT INTO websites(id,name,url,created_at) VALUES('a','Site','https://example.com','2026-11-19T00:00:00Z')");err!=nil{t.Fatal(err)}
 if _,err:=db.Exec("INSERT INTO websites(id,name,url,created_at) VALUES('b','Copy','https://example.com','2026-11-19T00:00:00Z')");err==nil{t.Fatal("duplicate URL accepted")}
 if _,err:=db.Exec("INSERT INTO checks(id,website_id,checked_at,success) VALUES('c','missing','2026-11-19T00:00:00Z',1)");err==nil{t.Fatal("invalid foreign key accepted")}
 if err:=Migrate(ctx,db);err!=nil{t.Fatal("repeated migration",err)}
 if err:=db.Close();err!=nil{t.Fatal(err)}
 db,err=Open(ctx,file);if err!=nil{t.Fatal(err)};defer db.Close()
 var count int
 if err:=db.QueryRow("SELECT count(*) FROM websites").Scan(&count);err!=nil||count!=1{t.Fatalf("persistence count=%d err=%v",count,err)}
 err=db.QueryRow("SELECT 1 FROM schema_migrations WHERE version=1").Scan(&count)
 if err!=nil && err!=sql.ErrNoRows{t.Fatal(err)}
}
