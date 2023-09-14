package main

import (
    "database/sql"
    "log"
    "net/http"
    "000-crud-web-applications/backend/routes"

    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"

)

func main() {
    // Initialize the SQLite database
    db, err := sql.Open("sqlite3", "database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create tables if they don't exist
    createTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL,
            password TEXT NOT NULL
        )
    `
    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatal(err)
    }

    router := mux.NewRouter()
    router = routes.SetRoutes(router)

    // Start the server
    log.Println("Server is running on :8080")
    http.ListenAndServe(":8080", router)
}
