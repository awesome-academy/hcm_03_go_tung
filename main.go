package main

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "postgres://username:password@localhost:5432/dbname?sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Your router setup, routes, etc.
}