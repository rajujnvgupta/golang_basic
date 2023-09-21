package main 

import (
  "database/sql"
  "fmt"
  _ "github.com/jackc/pgx/v5/stdlib"
)
const (
  host     = "localhost"
  port     = 5433
  user     = "postgres"
  password = "postgres"
  dbname   = "postgres"
)
func main() {
  postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", postgresqlDbInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("Established a successful connection!")

}
