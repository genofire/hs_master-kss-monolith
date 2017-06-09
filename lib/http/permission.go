// Package that provides the logic of the webserver
package http

import "net/http"

// Format of a function to bind it to the middleware handler
type HasPermission func(string, int) (bool, error)

// Function to evaluate the permission and implement an error handling
func PermissionHandler(h func(w http.ResponseWriter, r *http.Request), perm HasPermission, permission int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := r.Header.Get("session")
		ok, err := perm(session, permission)
		if err != nil {
			http.Error(w, err.Error(), http.StatusGatewayTimeout)
			return
		}
		if ok {
			h(w, r)
			return
		}
		http.Error(w, "Not allowed", http.StatusForbidden)

	}
}
