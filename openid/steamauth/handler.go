package steamauth

import "net/http"

// SteamLogin ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	opId := newOpenId(r)
	switch opId.mode() {
	case "":
		http.Redirect(w, r, opId.authUrl(), 301)
	case "cancel":
		w.Write([]byte("Authorization cancelled"))
	default:
		steamId, err := opId.validateAndGetId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Do whatever you want with steam id
		w.Write([]byte(steamId))
	}
}
