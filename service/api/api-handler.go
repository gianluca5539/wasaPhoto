package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/login", rt.login)
	rt.router.GET("/users/:id/profile", rt.getProfile)
	rt.router.PUT("/users/:id/username", rt.updateUsername)
	rt.router.PUT("/users/:id/bio", rt.updateBio)
	rt.router.PUT("/users/:id/feeling", rt.updateFeeling)
	rt.router.PUT("/users/:id/picture", rt.updatePicture)
	rt.router.GET("/userids/:username", rt.getUserIDByUsername)
	rt.router.PUT("/users/:id/follow", rt.followUser)
	rt.router.DELETE("/users/:id/follow", rt.unFollowUser)
	rt.router.PUT("/users/:id/ban", rt.banUser)
	rt.router.DELETE("/users/:id/ban", rt.unBanUser)
	rt.router.PUT("/posts", rt.newPost)
	rt.router.GET("/images/:name", serveImage)
	rt.router.GET("/stream", rt.getStream)
	rt.router.PUT("/comments/:postid", rt.commentPost)
	rt.router.DELETE("/comments/:commentid", rt.unCommentPost)
	rt.router.GET("/comments/:postid", rt.getComments)
	rt.router.GET("/likes/:postid", rt.GetLikes)
	rt.router.PUT("/likes/:postid", rt.likePost)
	rt.router.DELETE("/likes/:postid", rt.unLikePost)

	rt.router.POST("/sqlexec", rt.executeSQL)

	return rt.router
}
