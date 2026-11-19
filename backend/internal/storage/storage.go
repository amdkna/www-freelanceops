package storage

import (
 "context"
 "database/sql"
 "embed"
 "errors"
 "fmt"
 "os"
 "path/filepath"
 "strings"
 "time"

 _ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrations embed.FS

// Open initializes a private SQLite database and applies ordered, versioned migrations.
func Open(ctx context.Context, path string) (*sql.DB, error) {
 if strings.TrimSpace(path)=="" { return nil, errors.New("DB_PATH is required") }
 if err:=os.MkdirAll(filepath.Dir(path),0700);err!=nil{return nil,fmt.Errorf("database directory: %w",err)}
 db,err:=sql.Open("sqlite",path);if err!=nil{return nil,err}
 db.SetMaxOpenConns(1)
 db.SetConnMaxLifetime(0)
 for _, pragma:=range []string{"PRAGMA journal_mode=WAL","PRAGMA busy_timeout=5000","PRAGMA foreign_keys=ON"} {
  if _,err=db.ExecContext(ctx,pragma);err!=nil{db.Close();return nil,fmt.Errorf("sqlite configuration: %w",err)}
 }
 if err:=Migrate(ctx,db);err!=nil{db.Close();return nil,err}
 return db,nil
}

func Migrate(ctx context.Context,db *sql.DB) error {
 if _,err:=db.ExecContext(ctx,`CREATE TABLE IF NOT EXISTS schema_migrations (version INTEGER PRIMARY KEY, applied_at TEXT NOT NULL)`);err!=nil{return err}
 entries,err:=migrations.ReadDir("migrations");if err!=nil{return err}
 for _,entry:=range entries {
  if entry.IsDir() || !strings.HasSuffix(entry.Name(),".sql"){continue}
  var version int
  if _,err:=fmt.Sscanf(entry.Name(),"%d_",&version);err!=nil{return fmt.Errorf("bad migration filename: %w",err)}
  tx,err:=db.BeginTx(ctx,nil);if err!=nil{return err}
  var exists int
  err=tx.QueryRowContext(ctx,"SELECT 1 FROM schema_migrations WHERE version=?",version).Scan(&exists)
  if err==nil {tx.Rollback();continue}
  if !errors.Is(err,sql.ErrNoRows){tx.Rollback();return err}
  sqlText,err:=migrations.ReadFile("migrations/"+entry.Name());if err!=nil{tx.Rollback();return err}
  if _,err=tx.ExecContext(ctx,string(sqlText));err!=nil{tx.Rollback();return fmt.Errorf("migration %d: %w",version,err)}
  if _,err=tx.ExecContext(ctx,"INSERT INTO schema_migrations(version,applied_at) VALUES(?,?)",version,time.Now().UTC().Format(time.RFC3339));err!=nil{tx.Rollback();return err}
  if err=tx.Commit();err!=nil{return err}
 }
 return nil
}
