// Package http provides the
// logic of the webserver
package http

import "net/http"

type HasPermission func(string, int) (bool, error)

//Function to evaluate the permission and implent an error handling
// Input: http response writer w, pointer to htto request r, bool variable HasPermission perm, int variable permission (form)
func PermissionHandler(h func(w http.ResponseWriter, r *http.Request), perm HasPermission, permission int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
			return
		}
		ok, err := perm(session.Value, permission)
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
