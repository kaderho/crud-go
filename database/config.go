package database

import (
    "database/sql"
)
 
func Connect() *sql.DB {
    dbDriver := "mysql"
    dbUser := "kader"
    dbPass := "derka"
    dbName := "crud_go"
 
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}