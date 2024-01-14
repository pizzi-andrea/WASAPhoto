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

        async logout(){
            this.$axios.defaults.headers.common['Authorization'] = ''
            localStorage.removeItem('token')
            this.$router.push("/");

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
			<h1 class="h2">Ciao, {{profile.user.username ? profile.user.username: "anonimo"}}</h1>
            <div class="">
                    <p> Followers  {{profile.follower}}</p>
					<p> Following  {{profile.following}}</p>
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

			<div class="row d-flex flex-row justify-content-start justify-content-between" id="body">
				
                <div class="card-group">
                    <div v-for="post in myStream">
                                <Post :post="post"></Post>
                    </div>
                </div>
			</div>
			
			

			<div class="row" id="footer"></div>
		</div>

		
				
	</div>
	<hr>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
</style>

