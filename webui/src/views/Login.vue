<script>

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: "",
			errorInUsername: false,
			msgErrorData: " Errore stringa 'username' mal formattata. \n Un username valido deve avere una lunghezza compresa tra 3 e 16 caratteri di qualunque tipo."
		}
	},
	methods: {
		refresh() {
			this.loading = true;
			this.errormsg = null;
            this.errorInUsername=false
            this.username = ""
			localStorage.removeItem("token")
			
		},

		async login(){
            this.dataValidator()
			let response = null
			let token
			try {

				if (this.errorInUsername) {
					throw ("Username non valido.") 
				}
				response = await this.$axios.post("/session", {
					name: `'${this.username}'`
				});
				
				
				switch(response.status) {
					case 200:
					case 201:
						token = response.data.Value
						localStorage.setItem("token", token)
						// this.$axios.defaults.headers.Authorization = 'Bearer ' + token
						this.$router.push('/users/' + token + '/')
						
						console.log('/users/' + token + '/')
					break
					case 400:
						this.$route.push('/error/client')
						break
					case 500:
						this.$route.push('error/server')
						


				}
				this.loading = false;

				
				
			
			
			
		} catch (e) {
				this.errormsg = e.toString();
				return 
			}
		},

		dataValidator() { this.errorInUsername =  !( this.username.length >= 3 && this.username.length <= 16 && /^.*$/.test(this.username) )}


	},
	mounted() {
		this.refresh()
	}
}



</script>

<template>
	
	<title>WasaPhoto - login</title>
	
	<div class = "container form-group  mr-4 mt-4">
			
			

				<div class="row">
					
					<div class="container">

						<div class="row  flex-row d-inline-flex">
							
							<h4>Username</h4>
							<div class="col-sm m-1 d-sm-inline-flex">
								
								<input v-model="username" id="usernameField" placeholder="username" minlength="3" maxlength="16"  class =" m-1 form-control" @input ="dataValidator" @focus="refresh">
								<button class="btn btn-primary m-4" @click="login">Login/Sigin</button>
							</div>
							
						</div>
						
						
					</div>

				</div>
				
		
	</div>

		
		<hr>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<br>
		<ErrorMsg v-if="errorInUsername" :msg="msgErrorData"></ErrorMsg>		
	
</template>

<style>
</style>

