package utils

import (
    "net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        username, password, _ := r.BasicAuth()

        if !Authenticate(username, password) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}

/*
func Authenticate(username, password string) bool {
    // Hardcoded credentials for demonstration
    return username == "admin" && password == "password"
}
*/