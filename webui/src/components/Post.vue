<script>
export default {
	props: ['post'],

  data: function() {
		return {
      errormsg: "",
      username: "",
      myUsername: "",
      txtComment: "",
      users_likes: [],
      likes: 0,
      img: this.$axios.defaults.baseURL + "/images/" + this.post.refer,
      like: false,
      commentPerform: false,
      loading: false,
      comments: [],



    }
  },

  methods: {
        async refresh() {},

        async getMyUsername(){
          let response = {}
          try {
            response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/")
          }catch(e){
            this.errormsg = e.toString()
          }

          switch(response.status){
            case 200:
              this.myUsername =  response.data.user.username
              break;
            case 400:
              this.$router.push("/error/client");
              break;
            case 500:
              this.$router.push("/error/server");
              break;
            
          }

        },
        async delLike(){
          try{
            let response = {}
            response = await this.$axios.delete("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/" + localStorage.getItem('token'))
            console.log(response)
            
            switch(response.status){
              case 204:
                this.like = false;
                this.countLikes();
                break;
              case 404:
                //TODO
                break;
              case 400:
                //todo
                break;
              case 500:
                //todo
                break;
              default:
                //todo
                
            }
          }catch(e){
            this.errormsg = e.toString()

          }
        
        },
        async putLike(){
          try{
            let response = {}
            response = await this.$axios.put("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/" + localStorage.getItem('token'))
            switch(response.status){
              case 201:
                this.like = true;
                this.countLikes();
                break;
              case 204:
                break;
              case 404:
                //TODO
                break;
              case 400:
                //todo
                break;
              case 500:
                //todo
                break;
              default:
                //todo
                
            }
          }catch(e){
            this.errormsg = e.toString()
          }

        },

        async getImage(){
          let response = {}
          try {
            response = await this.$axios.get("/images/" + this.post.refer);
          
          }catch(e){
            this.errormsg = e.toString()
          }
          switch(response.status){
            case 200:
              break;
            case 400:
              this.$router.push("/error/client");
              break;
            case 500:
              this.$router.push("/error/server");
              break;
          }
        },

        async putComment(){
          try{
            let response = {}
            response = await this.$axios.post("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/comments/", {
              author: {
                username: this.myUsername
              },
              text: this.txtComment,

            });
            switch(response.status){
              case 201:
                this.disableComment();
                this.refreshComments();
                break;
              case 404:
                //TODO
                break;
              case 400:
                //todo
                break;
              case 500:
                //todo
                break;
              default:
                //todo
                
            }
          }catch(e){
            this.errormsg = e.toString()
          }
        },

        enableComment(){
          this.commentPerform = true;
        },

        disableComment(){
          this.commentPerform = false;
          this.txtComment = ""
        },

        clearTextComment(){
          this.txtComment = ""
        },

        async getLikes(){
          try{
            let response = {}
            response = await this.$axios.get("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/")
            console.log(response)
            
            switch(response.status){
              case 200:
                this.users_likes = response.data;
                break;
              case 204:
                this.users_likes = [];
                break;
              case 404:
                //TODO
                break;
              case 400:
                //todo
                break;
              case 500:
                //todo
                break;
              default:
                //todo
                
            }

            

          }catch(e){
            this.errormsg = e.toString()
          }
        },

        
        async getUsername(){
          let response = {}
          try {
            response = await this.$axios.get("/users/" + this.post.owner + "/")
          }catch(e){
            this.errormsg = e.toString()
          }

          switch(response.status){
            case 200:
              this.username =  response.data.user.username
              break;
            case 400:
              this.$router.push("/error/client");
              break;
            case 500:
              this.$router.push("/error/server");
              break;
            
          }
        },

        async showProfile(id){
          return 
        },

        async countLikes(){
          await this.getLikes()
          this.likes = this.users_likes.length

        },


        async isLiked(){
          try{
            let response = {}
            response = await this.$axios.get("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/" + localStorage.getItem('token'));
            
            switch(response.status){
              case 200:
                this.like = true
                break;
              case 204:
                this.like = false;
              default:
                //todo
            }
          }catch(e){

          }
        },

        async refreshComments(){
          try{
            let response = {}
            response = await this.$axios.get("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/comments/");
            
            switch(response.status){
              case 200:
                this.comments = response.data
                break;
              case 204:
                this.comments = []
              default:
                //todo
            }
          }catch(e){

          }

        }


        
      
        
    },
    mounted(){
        this.getUsername();
        this.isLiked();
        this.getImage();
        this.getMyUsername();

        this.users_likes = this.post.likes;
        this.likes = this.users_likes ? this.users_likes.length : 0;
        this.comments = this.post.comments;
    },
}
</script>

<template>
	<div class="card m-4" style="width: 18rem;" id={{post.refer}}>
    <div class="card-header">
      <a @click=" showProfile(post.owner)"> <i>Postata da {{ username }}</i> </a>
      
    </div>
    <img class="card-img-top img-fluid object-fit-xxl-contain border rounded" :src="img" alt="img">
  <div class="card-body">
    <p class="card-text"> {{post.descriptionImg}} </p>
    
    <!-- -->
    <button type="button" @click="delLike" class="btn btn-danger" v-if="like">
      <span class="badge badge-light"> <img src="./icons/lost-like.svg"> </span>  <span class="badge badge-light">{{likes}}</span>
    </button>
    <button type="button" @click="putLike" class="btn btn-primary" v-else>
      <span class="badge badge-light"> <img src="./icons/need-like.svg"> </span> <span class="badge badge-light">{{likes}}</span>
    </button>


    <p class="card-text"><small class="text-muted">Aggiunta il: {{ post.timeUpdate }}</small></p>
    
    
    <div class="list-group">
      <a @click="showProfile(comment.author)"  class="list-group-item list-group-item-action flex-column align-items-start" v-for="comment in comments" :key="comment.commentId">
        <div class="d-flex w-100 justify-content-between">
          <h5 class="mb-1"></h5>
          <small>{{comment.timeStamp}}</small>
        </div>
        <p class="mb-1">{{comment.text}}</p>
          <small></small>
      </a>
    
      <div class="form-floating m-2" v-if="commentPerform">
        <textarea class="form-control" placeholder="Leave a comment here" id="floatingTextarea2" style="height: 100px" v-model="txtComment"></textarea>
        <label for="floatingTextarea2">Comments</label>

        <div class="btn-group" role="group m-4 p-2 d-flex">
          <button type="button" class="btn btn-primary" @click="putComment">Invia</button>
          <button type="button" class="btn btn-secondary" @click="clearTextComment">Cancella</button>
          <button type="button" class="btn btn-secondary"> <span class="badge badge-light" @click="disableComment"> <img src="./icons/cancel.svg"> </span> </button>
        </div>

      </div>
      <button type="button" @click="enableComment"  class="btn btn-light m-2" v-else>
        <small> commenta </small> <span class="badge badge-light"><img src="./icons/write.svg"> </span>
      </button>
     
    </div>


    <div class="spinner-border" role="status" v-if="loading">
      <span class="visually-hidden">Loading...</span>
    </div>
    
    
  </div>
</div>
<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
</style>