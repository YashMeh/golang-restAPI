package store

import (
	"net/http"

	"github.com/gorilla/mux"
)

// var controller = &Controller{Repository: Repository{}}
var controller = &Controller{}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"GetBooks",
		"GET",
		"/api/books",
		controller.getBooks,
	},
	Route{
		"GetBook",
		"GET",
		"/api/book/{id}",
		controller.getBook,
	},
	Route{
		"PostBook",
		"POST",
		"/api/books",
		controller.createBook,
	},
	Route{
		"UpdateBook",
		"PUT",
		"/api/book/{id}",
		controller.updateBook,
	},
	Route{
		"DeleteBook",
		"DELETE",
		"/api/book/{id}",
		controller.deleteBook,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		//can be used to pass custom middlewares(like logging)
		//handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
