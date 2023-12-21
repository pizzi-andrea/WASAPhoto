<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			token: "",
			username: "",
			errorInUsername: true,
			msgErrorData: " Errore stringa 'username' mal formattata. \n Un username valido deve avere una lunghezza compresa tra 3 e 16 caratteri di qualunque tipo."
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			
		},

		async login(){
			try {

				if (this.errorInUsername) {
					throw new TypeError("Username non valido.") 
				}
				let response = await this.$axios.post("/session", {
					name: `'${this.username}'`
				});
				this.token = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			
		},

		dataValidator() { this.errorInUsername =  !( this.username.length >= 3 && this.username.length <= 16 && /^.*$/.test(this.username) )}


	},
	mounted() {
		this.refresh()
	}
}



</script>

<template>
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

		<div>
			<h5>Username</h5>
			<input v-model="username" placeholder="write your username hear" minlength="3" maxlength="16"  @input ="dataValidator">

		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="errorInUsername" :msg="msgErrorData"></ErrorMsg>
		
		



		
	</div>
</template>

<style>
</style>

