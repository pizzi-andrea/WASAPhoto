<script>
import token from "../router/index"
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
			
		},

		async login(){
            this.dataValidator()
			try {

				if (this.errorInUsername) {
					throw ("Username non valido.") 
				}
				let response = await this.$axios.post("/session", {
					name: `'${this.username}'`
				});
				token = response.data;
			
			this.loading = false;
			
			
			this.$router.push('/users/' + this.token.Value + '/')
			
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
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Login</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="login">
						Sig-in/Log-in
					</button>
				</div>
			</div>
		</div>

		<div class=" align-items-center ">
			<h5>Username</h5>
			
			<input v-model="username" placeholder="" minlength="3" maxlength="16"  @input ="dataValidator" @focus="refresh">

		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="errorInUsername" :msg="msgErrorData"></ErrorMsg>		
	</div>
</template>

<style>
</style>

