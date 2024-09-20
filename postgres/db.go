package postgres

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type Storage interface{
    Create(*Type) error 
    GetOne(id string)(*Type, error)
    // yada yada you know the drill....
}


type Postgresql struct{
    db *sql.DB
}

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
