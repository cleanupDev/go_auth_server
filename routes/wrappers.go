package routes

import "net/http"

func methodNotAllowedWrapper(handler http.HandlerFunc, allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range allowedMethods {
			if r.Method == method {
				handler(w, r)
				return
			}
		}

		w.Header().Set("Allow", joinStrings(allowedMethods, ", "))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func joinStrings(strings []string, delimiter string) string {
	if len(strings) == 0 {
		return ""
	}
	result := strings[0]
	for i := 1; i < len(strings); i++ {
		result += delimiter + strings[i]
	}
	return result
}
