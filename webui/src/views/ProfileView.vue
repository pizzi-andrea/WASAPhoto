<script>
import Profile from "../components/Profile.js"
export default {
	
    data: function () {
        return {
            errormsg: null,
            loading: false,
            profile: Profile,
			myStream: []
            
        };
    },
    methods: {
        async refresh() {
            let response;
            try {
                response = await this.$axios.get(this.$route.path);
            }
            catch (e) {
                return;
            }
            switch (response.status) {
                case 200:
                    this.profile = new Profile();
                    this.profile = response.data;
                    break;
                case 404:
                    this.$router.push("error/NotFound");
                    break
                case 500:
                    this.$router.push("error/ServerError");
                    break
                default:
                    this.$router.push("error/ServerError");
            }

			try {
                response = await this.$axios.get(this.$route.path + "myStream/");
            }
            catch (e) {
                return;
            }
            switch (response.status) {
                case 200:
				case 204:
                    this.myStream = response.data;
                    break;
                case 404:
                    this.$router.push("error/NotFound");
                    break
                case 500:
                    this.$router.push("error/server");
                    break
                default:
                    this.$router.push("error/server");
            }
        },

        logout(){
            this.$axios.defaults.headers.common['Authorization'] = ''
            localStorage.clear()
            this.$router.replace("/");

        },

        async uploadUsername(username){
            
            try{
                await this.$axios.put(this.$route.path, {
                    username: username
                });
            }catch(e){

            }
        }
    },
    mounted() {
        this.refresh();
    },

    

}



</script>

<template>
	<title>WasaPhoto</title>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Benvenuto {{ profile.user.username }}</h1>
            <div class="vr m-4"></div>
            <div class="vstack gap-2 col-md-5 mx-auto">

                    <h3> Followers</h3>  
                    <h5>{{profile.follower}}</h5>
					<h3> Following</h3> 
                    <h5> {{profile.following}}</h5>
                </div>
			<div class="btn-toolbar mb-2 mb-md-0">

                
				
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="refresh">
						Aggiorna stream
					</button>
                    <button type="button" class="btn btn-sm btn-warning" @click="uploadUsername">
						Aggiorna username
					</button>
					<button type="button" class="btn btn-sm btn btn-danger" @click="logout">
						Logout
					</button>
				</div>
			</div>
		</div>

		<div class = 'container float-start'>
			<div class="row" id="top">
			</div>

			<div class="row d-inline p-2 flex-row justify-content-start justify-content-between col-3 position-absolute top-25 start-50 translate-middle-x" id="body">
				
                <div class="">
                    <div v-for="post in myStream">
                                <Post :post="post"></Post>
                    </div>
                </div>
			</div>
			
			

			<div class="row" id="footer"></div>
		</div>

		
				
	</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
</style>

