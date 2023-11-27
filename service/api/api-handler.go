package api

import (
	"net/http"
)

/* resources */

const r_users = "/users/"                              // this resource rappresent collection of users
const r_user = r_users + ":uid/"                       // Resource rappresent a single user
const r_followers = r_user + "followers/"              // this resurce rappresent the followers of specific user
const r_follower = "/users/:uid/followers/:followerId" // this resource rappresent follower
const r_login = "/session/"
const r_userlog = r_login + ":tokenId" // this endpont rappresent logged user
const r_root = "/"

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	//
	rt.router.POST(r_login, rt.doLogin)

	//list registred users
	rt.router.GET(r_users, rt.listUsers)

	// assign or update username
	rt.router.PUT(r_user, rt.setMyUserName)

	// get specific user profile
	rt.router.GET(r_user, rt.getUserProfile)

	// Logout User
	rt.router.DELETE(r_userlog, rt.doLogout)

	/*
		// get all followers
		rt.router.GET("/users/:uid/followers", rt.listFollowers)

		// unfollow user
		rt.router.DELETE("/users/:uid/followers/:followerId", rt.unfollowUser)

		//follow user
		rt.router.PUT("/users/:uid/followers/:followerId", rt.followUser)

		// get following users
		rt.router.GET("/users/:uid/following", rt.getFollowed)

		// get followed user
		rt.router.GET("/users/{uid}/following/:followingId", rt.getFollowing)

		// list personal stream photos
		rt.router.GET("/users/:uid/myStream", rt.getMyStream)

		// get photo from stream
		rt.router.GET("/users/:uid/myStream/:photoId", rt.getPhotoMyStream)

		// banned users
		rt.router.GET("/users/:uid/banned/", rt.listBannedUser)

		// ban user identificated by *uid*
		rt.router.PUT("/users/:uid/banned/:bannedId", rt.banUser)

		// unban user identificated by uid
		rt.router.DELETE("/users/:uid/banned/:bannedId", rt.unbanUser)

		// update photo
		rt.router.POST("/users/:uid/myPhotos", rt.uploadPhoto)

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

		// Logs in the user
		rt.router.POST("/session", rt.doLogin)

		// Special routes
		rt.router.GET("/liveness", rt.liveness)
	*/

	// test method
	rt.router.GET(r_root, rt.hello)
	return rt.router

}
