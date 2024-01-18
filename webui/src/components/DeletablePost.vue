<script>
import { RouterLink, RouterView } from 'vue-router'
export default {
	props: ['post'],

  data: function() {
		return {
      errormsg: "",
      username: "",
      myUsername: "",
      users_likes: [],
      likes: 0,
      img: this.$axios.defaults.baseURL + "/images/" + this.post.refer,
      loading: false,
      comments: [],
      i: localStorage.getItem('token')



    }
  },

  methods: {
        async refresh() {

          this.refreshComments();
          this.getUsername();
          this.getImage();
          this.getMyUsername();
        },

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

        

        async countLikes(){
          await this.getLikes()
          this.likes = this.users_likes ? this.users_likes.length : 0

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

        },

        async deletePhoto(){
            try{
            let response = {}
            response = await this.$axios.delete("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/");
            console.log(response)
            switch(response.status){
              case 204:
                document.getElementById(post.refer).remove();
                break;
              case 400:
                break;
              case 404:
                break;
              case 500:
                break;
              default:
                //todo
            }
          }catch(e){

          }

        }


        
      
        
    },
    mounted(){
        this.refresh();

        this.users_likes = this.post.likes;
        this.likes = this.users_likes ? this.users_likes.length : 0;
        this.comments = this.post.comments;
    },
}
</script>

<template>
	<div class="card m-4" style="" :id="post.refer">
    <div class="card-header">

      <RouterLink :to=" '/users/' + post.owner + '/profile' ">
        <a> <i>Postata da {{ username }}</i> </a>
      </RouterLink>
      
      
    </div>
    <img class="card-img-top img-fluid object-fit-xxl-contain border rounded" :src="img" alt="img">
  <div class="card-body">
    <p class="card-text"> {{post.descriptionImg}} </p>
    
    <!-- -->
    <button type="button" @click="deletePhoto" class="btn btn-danger">
      <span class="badge badge-light"> <img src="./icons/trash.svg"> </span>  <span class="badge badge-light">{{likes}}</span>
    </button>
   


    <p class="card-text"><small class="text-muted">Aggiunta il: {{ post.timeUpdate }}</small></p>
    
    
    <div class="list-group">
      <RouterLink :to=  "comment.author.uid == i ? '#' : '/users/' + comment.author.uid + '/profile' " v-for="comment in comments" :id="comment.commentId">
        <a @click="showProfile(comment.author.uid)"  class="list-group-item list-group-item-action flex-column align-items-start border rounded" >
          <div class="d-flex w-100 justify-content-between">
            <h5 class="mb-1"></h5>
            <small><b><i>{{ comment.author.username}}</i></b></small>
            <small>{{comment.timeStamp}}</small>
          </div>
          <p class="mb-1">{{comment.text}}</p>
            <small></small>
        </a>
    </RouterLink>
     
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