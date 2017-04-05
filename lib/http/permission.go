package http

import "net/http"

type HasPermission func(string, int) (bool, error)

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
