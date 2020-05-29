package middlew

import (
	"net/http"

	"github.com/DanteCV2/goltweet/db"
)

/*CheckDB check the status in DB*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConn() == 0 {
			http.Error(w, "Lost conection with DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
