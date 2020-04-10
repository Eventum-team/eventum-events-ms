package service
import (
	u "ev-events-ms/utils"
	"net/http"
)

var NotFoundHandler = func() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u.Respond(w, u.Message(false, "This resources was not found on our server"))
	})
}
