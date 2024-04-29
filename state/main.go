package main

import (
    "net/http"
    "github.com/gorilla/sessions"
	"fmt"
)

var (
    store = sessions.NewCookieStore([]byte("secret-key"))
)

func main() {
    http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    
    // Check if user is authenticated
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "You are not authenticated", http.StatusForbidden)
        return
    }
    
    // Print welcome message
    fmt.Fprintln(w, "Welcome to the secure area!")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    
    // Authentication goes here
    // For now, we'll assume authentication is successful
    
    session.Values["authenticated"] = true
    session.Save(r, w)
    fmt.Fprintln(w, "You are logged in.")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    session.Values["authenticated"] = false
    session.Save(r, w)
    fmt.Fprintln(w, "You are logged out.")
}
