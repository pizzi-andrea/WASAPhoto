<script>
export default {
	props: [],

  data: function() {
		return {
      errormsg: "",
      img: null,
      description: "",
      
     



    }
  },

  methods: {
        
        async uploadPost(){
          const form =  new FormData();
          form.append("img", this.img);
          form.append("desc", this.description);
          let response = null 
          try{
            
            response = await this.$axios.post("/users/" + localStorage.getItem('token')  + "/myPhotos/", form);

            switch(response.status){
              case 201:
                break;
              
            }
          }catch(e){
            console.log(e);
                switch (e.response.status) {
                    case 400:
                        $router.push("/error/400");
                        break;
                    case 401:
                        this.$router.push("/error/401");
                        break;
                    case 403:
                        this.$router.push("/error/403");
                        break;
                    case 404:
                        $router.push("error/404");
                        break;
                    case 500:
                        this.$router.push("/error/500");
                        break;
                }
          }
        },

        showImg(event){
          const selectedFile = event.target.files[0];
          if (selectedFile) {
            this.img = selectedFile;
          }

        },

        

        },
    mounted(){
      

    },
}
</script>

<template>
	<div class="card m-4" style="width: 40rem;">
    <div class="card-header">
      <small>Postato da</small> 
      
    </div>
    <form>
    <div class="card-img-top img-fluid object-fit-xxl-contain border rounded"></div>
  
    <div class="card-body">
      
      <div class="input-group mb-3">
        <span class="input-group-text mb-3">descrizione</span>
          <textarea class="form-control " aria-label="descrizione" v-model="description" resize="none"></textarea>
      </div>

      <div class="input-group mb-3">
        <input type="file" class="form-control" id="inputGroupFile04" aria-describedby="inputGroupFileAddon04" aria-label="Upload" v-on:change="showImg">
        <button class="btn btn-outline-secondary" type="button" id="inputGroupFileAddon04" @click="uploadPost">Button</button>
      </div>
      
      


      <p class="card-text"><small class="text-muted">Aggiunta il </small></p>
      
      
      <div class="list-group">
        
            <small><i>commenti</i></small>
      </div>
    </div>
    </form>

    
</div>
<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
</style>
