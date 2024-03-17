/*
 * SystemK
 *
 * API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
 *
 * API version: 1.0.0
 */
package mware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateSubDivisionV1",
		strings.ToUpper("Post"),
		"/v1/subdivision",
		CreateSubDivisionV1,
	},

	Route{
		"CreateUserV1",
		strings.ToUpper("Post"),
		"/v1/user",
		CreateUserV1,
	},

	Route{
		"CreateWorkPositionV1",
		strings.ToUpper("Post"),
		"/v1/workposition",
		CreateWorkPositionV1,
	},

	Route{
		"CreateWorkerV1",
		strings.ToUpper("Post"),
		"/v1/worker",
		CreateWorkerV1,
	},

	Route{
		"DeleteSubDivisionV1",
		strings.ToUpper("Delete"),
		"/v1/subdivision",
		DeleteSubDivisionV1,
	},

	Route{
		"DeleteUserV1",
		strings.ToUpper("Delete"),
		"/v1/user",
		DeleteUserV1,
	},

	Route{
		"DeleteWorkPositionV1",
		strings.ToUpper("Delete"),
		"/v1/workposition",
		DeleteWorkPositionV1,
	},

	Route{
		"GetSubDivisionV1",
		strings.ToUpper("Get"),
		"/v1/subdivision",
		GetSubDivisionV1,
	},

	Route{
		"GetUserV1",
		strings.ToUpper("Get"),
		"/v1/user",
		GetUserV1,
	},

	Route{
		"GetWorkPositionV1",
		strings.ToUpper("Get"),
		"/v1/workposition",
		GetWorkPositionV1,
	},

	Route{
		"GetWorkerV1",
		strings.ToUpper("Get"),
		"/v1/worker",
		GetWorkerV1,
	},

	Route{
		"UpdateSubDivisionV1",
		strings.ToUpper("Put"),
		"/v1/subdivision",
		UpdateSubDivisionV1,
	},

	Route{
		"UpdateUserV1",
		strings.ToUpper("Put"),
		"/v1/user",
		UpdateUserV1,
	},

	Route{
		"UpdateWorkPositionV1",
		strings.ToUpper("Put"),
		"/v1/workposition",
		UpdateWorkPositionV1,
	},

	Route{
		"UpdateWorkerV1",
		strings.ToUpper("Put"),
		"/v1/worker",
		UpdateWorkerV1,
	},
}
