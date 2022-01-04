package routes

import "net/http"

var loginRoute = []Route{
	{
		URI:         "/login",
		Method:      http.MethodPost,
		Function:    func(rw http.ResponseWriter, r *http.Request) {},
		RequireAuth: false,
	},
}
