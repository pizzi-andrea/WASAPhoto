package api

import (
	"net/http"
)

/* resources */
const r_root = "/"
const r_users = r_root + "users/"              // Endpoint rappresent the list of WasaPhoto users
const r_user = r_users + ":uid/"               // Endpont rappresent a WasaPhoto user
const r_followers = r_user + "followers/"      // Endpoint rappresent the followers of specific user
const r_follower = r_followers + ":followerId" // Endpont rappresent follower a user
const r_login = r_root + "session"
const r_banned = r_user + "banned/"
const r_userBanned = r_banned + ":bannedId"

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	rt.router.GET(r_root, rt.hello)

	//Logs in the user
	rt.router.POST(r_login, rt.doLogin)

	//list registred users
	rt.router.GET(r_users, rt.listUsers)

	// assign or update username
	rt.router.PUT(r_user, rt.setMyUserName)

	// get specific user profile
	rt.router.GET(r_user, rt.getUserProfile)

	// get all followers
	rt.router.GET(r_followers, rt.listFollowers)

	// unfollow user
	rt.router.DELETE(r_follower, rt.unfollowUser)

	//follow user
	rt.router.PUT(r_follower, rt.followUser)

	/*

		// get followed user
		rt.router.GET("/users/{uid}/following/:followingId", rt.getFollowing)
	*/
	/*

		// list personal stream photos
		rt.router.GET("/users/:uid/myStream", rt.getMyStream)

		// get photo from stream
		rt.router.GET("/users/:uid/myStream/:photoId", rt.getPhotoMyStream)
	*/

	// banned users
	rt.router.GET(r_banned, rt.listBannedUsers)

	// ban user identificated by *uid*
	rt.router.PUT(r_userBanned, rt.banUser)

	// unban user identificated by uid
	rt.router.DELETE(r_userBanned, rt.unbanUser)

	// update photo
	rt.router.POST("/users/:uid/myPhotos", rt.uploadPhoto)
	/*
		// list stream photos updated
		rt.router.GET("/users/:uid/myPhotos", rt.listPhoto)

		// delete photo updated
		rt.router.DELETE("/users/:uid/myPhotos/:photoId/", rt.deletePhoto)

		// get photo
		rt.router.GET("/users/:uid/myPhotos/:photoId/", rt.getPhoto)

		// get likes collected by photo
		rt.router.GET("/users/:uid/myPhotos/:photoId/likes/", rt.getLikes)

		// put like a photo
		rt.router.PUT("/users/:uid/myPhotos/:photoId/likes/:likeUserId:", rt.likePhoto)

		// remove like a photo
		rt.router.DELETE("/users/:uid/myPhotos/:photoId/likes/:likeUserId:", rt.unlikePhoto)

		// add comment a photo
		rt.router.POST("/users/:uid/myPhotos/:photoId/comments/", rt.commentPhoto)

		// get comments on photo
		rt.router.GET("/users/:uid/myPhotos/:photoId/comments/", rt.getComments)

		// delete comment on photo
		rt.router.DELETE("/users/:uid/myPhotos/:photoId/comments/:commentId", rt.uncommentPhoto)

		// get comment on photo
		rt.router.GET("/users/:uid/myPhotos/:photoId/comments/:commentId", rt.getComment)


		// Special routes
		rt.router.GET("/liveness", rt.liveness)
	*/
	return rt.router

}
