package postgres

import (
    "database/sql"
)


func ConnectDB() *sql.DB {
    connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    
    return db
}
