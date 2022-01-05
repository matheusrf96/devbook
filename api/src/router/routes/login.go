package routes

import (
	"api/src/router/controllers"
	"net/http"
)

var loginRoute = []Route{
	{
		URI:         "/login",
		Method:      http.MethodPost,
		Function:    controllers.Login,
		RequireAuth: false,
	},
}
