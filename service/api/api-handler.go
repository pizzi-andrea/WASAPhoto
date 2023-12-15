package api

import (
	"net/http"
)

// endpoints path
const (
	r_root        = "/"                         // "/"
	r_users       = r_root + "users/"           // "/users/"
	r_user        = r_users + ":uid/"           // "/users/{uid}/"
	r_followers   = r_user + "followers/"       // "/users/{uid}/followers/"
	r_follower    = r_followers + ":followerId" // "/users/{uid}/followers/{followerId}"
	r_login       = r_root + "session"          // "/session
	r_banned      = r_user + "banned/"          // "/users/{uid}/banned/"
	r_userBanned  = r_banned + ":bannedId"      // "/users/{uid}/banned/{bannedId}"
	r_followed    = r_user + "followed"         // "/users/{uid}/followed"
	r_myPhotos    = r_user + "myPhotos/"        // "/users/{uid}/myPhotos/"
	r_myPhoto     = r_myPhotos + ":photoId/"    // "/users/{uid}/myPhotos/{photoId}/"
	r_myStream    = r_user + "myStream/"        // "/users/{uid}/myStream/"
	r_streamPhoto = r_myStream + ":photoId"     // "/users/{uid}/myStream/{photoId}"
	r_comments    = r_myPhoto + "comments/"     // "/users/{uid}/myPhotos/{photoId}/comments/"
	r_comment     = r_comments + ":commentId"   // "/users/{uid}/myPhotos/{photoId}/comments/{commentId}"
	r_likes       = r_myPhoto + "likes/"        // "/users/{uid}/myPhotos/{photoId}/likes/"
	r_like        = r_likes + ":likeUserId"     // "/users/{uid}/myPhotos/{photoId}/likes/{likeUserId}"
	r_image       = r_root + "images/:photoId"  // "/images/{photoId}"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//   Register routes

	rt.router.GET(r_root, rt.wrap(rt.getContextReply))

	//  Logs in the user
	rt.router.POST(r_login, rt.wrap(rt.doLogin)) // required

	//  list registred users
	rt.router.GET(r_users, rt.wrap(rt.listUsers))

	//   assign or update username
	rt.router.PUT(r_user, rt.wrap(rt.setMyUserName)) // required

	//   get specific user profile
	rt.router.GET(r_user, rt.wrap(rt.getUserProfile)) // required

	//   get all followers
	rt.router.GET(r_followers, rt.wrap(rt.listFollowers))

	//   unfollow user
	rt.router.DELETE(r_follower, rt.wrap(rt.unfollowUser)) // required

	//  follow user
	rt.router.PUT(r_follower, rt.wrap(rt.followUser)) // required

	//   get followed user
	rt.router.GET(r_followed, rt.wrap(rt.getFollowed))

	//   list personal stream photos
	rt.router.GET(r_myStream, rt.wrap(rt.getMyStream)) // required

	//   banned users
	rt.router.GET(r_banned, rt.wrap(rt.listBannedUsers)) // fixed

	//   ban user identificated by *uid*
	rt.router.PUT(r_userBanned, rt.wrap(rt.banUser)) // required

	//   unban user identificated by uid
	rt.router.DELETE(r_userBanned, rt.wrap(rt.unbanUser)) // required

	//   update photo
	rt.router.POST(r_myPhotos, rt.wrap(rt.uploadPhoto)) // required

	//   list stream photos updated
	rt.router.GET(r_myPhotos, rt.wrap(rt.listPost))

	//   delete photo updated
	rt.router.DELETE(r_myPhoto, rt.wrap(rt.deletePhoto)) // required

	//   get photo
	rt.router.GET(r_myPhoto, rt.wrap(rt.getPost))

	//   get likes collected by photo
	rt.router.GET(r_likes, rt.wrap(rt.getLikes))

	//   put like a photo
	rt.router.PUT(r_like, rt.wrap(rt.likePhoto)) // required

	//   remove like a photo
	rt.router.DELETE(r_like, rt.wrap(rt.unlikePhoto)) // required

	//  add comment a photo
	rt.router.POST(r_comments, rt.wrap(rt.commentPhoto)) // required

	//  get comments on photo
	rt.router.GET(r_comments, rt.wrap(rt.getComments))

	//  delete comment on photo
	rt.router.DELETE(r_comment, rt.wrap(rt.uncommentPhoto)) // required

	//  get comment on photo
	rt.router.GET(r_comment, rt.wrap(rt.getComment))

	// specific endpoint used for get photo data
	rt.router.GET(r_image, rt.wrap(rt.getImage))
	//  Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router

}
