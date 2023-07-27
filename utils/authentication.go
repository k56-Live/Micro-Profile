package utils

import (
        "fmt"
        "net/http"
)

// AuthenticateUser is a middleware function to check if the user is authenticated
func AuthenticateUser(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                // Implement your authentication logic here
                // For example, you can check if the user is logged in or has a valid session token
                // If the user is not authenticated, redirect them to the login page or return an error response

                // Example implementation (dummy logic)
                isAuthenticated := true
                if !isAuthenticated {
                        http.Redirect(w, r, "/login", http.StatusSeeOther)
                        return
                }

                // If the user is authenticated, call the next handler in the chain
                next.ServeHTTP(w, r)
        })
}
