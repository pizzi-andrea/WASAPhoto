<script>
export default {
	props: ['post'],

  data: function() {
		return {
      errormsg: "",
      username: "",
      likes: [],
      img: "",
      like: false

    }
  },

  methods: {
        async refresh() {
          this.getUsername()
          this.getLikes()
          this.getImage()
        },
        async putLike(){

        },

        async getImage(){
          this.img = ""
          let response = {}
          try {
            response = await this.$axios.get("/images/" + this.post.refer);
          
          }catch(e){
            console.log( e.toString());
				    return

          }
          switch(response.status){
            case 200:
              this.img =  response.data
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

        },

        async getLikes(){
          let response = {}
          try {
            response = await this.$axios.get("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/")
          
          }catch(e){
            console.log( e.toString());
				    return

          }
          switch(response.status){
            case 200:
              this.likes =  response.data ? response.data : []
              break;
            case 400:
              this.$router.push("/error/client");
              break;
            case 500:
              this.$router.push("/error/server");
              break;
          }
          
        },

        async getUsername(){
          let response = {}
          try {
            response = await this.$axios.get("/users/" + this.post.owner + "/")
          }catch(e){
            console.log(e.toString());
				    return

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
        
        
    },
    mounted(){
          this.refresh();    
    },
}
</script>

<template>
	<div class="card m-4" style="width: 18rem;" id={{post.refer}}>
    <div class="card-header">
      <a @click=" showProfile(post.owner)"> <i>Postata da {{ username }}</i> </a>
      
    </div>
    <img class="card-img-top img-fluid object-fit-xxl-contain border rounded" :src=img alt="img">
  <div class="card-body">
    <p class="card-text">{{post.descriptionImg}}</p>
    
    <!-- -->
    <button type="button" @click="putLike" class="btn btn-danger" v-if="like">
      <span class="badge badge-light"> <img src="./icons/lost-like.svg"> </span>  <span class="badge badge-light">{{likes.length}}</span>
    </button>
    <button type="button" @click="putLike" class="btn btn-primary" v-else>
      <span class="badge badge-light"> <img src="./icons/need-like.svg"> </span> <span class="badge badge-light">{{likes.length}}</span>
    </button>

    <p class="card-text"><small class="text-muted">Aggiunta il: {{ post.timeUpdate }}</small></p>
    
    <div class="list-group overflow-auto" >
      <a @click="showProfile(comment.author)"  class="list-group-item list-group-item-action flex-column align-items-start" v-for="comment in post.comments" :key="comment.commentId">
        <div class="d-flex w-100 justify-content-between">
          <h5 class="mb-1"></h5>
          <small>{{comment.timeStamp}}</small>
        </div>
        <p class="mb-1">{{comment.text}}</p>
        <small></small>
      </a>

      <button type="button" @click="putLike" class="btn btn-light m-2">
        <small> commenta </small> <span class="badge badge-light"><img src="./icons/write.svg"> </span>
      </button>
    </div>
    
    
  </div>
</div>
</template>

<style>
</style>