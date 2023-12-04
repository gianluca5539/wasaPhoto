package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/login", rt.login)
	rt.router.GET("/users/:id/profile", rt.getProfile)
	rt.router.POST("/sqlexec", rt.executeSQL)

	return rt.router
}
