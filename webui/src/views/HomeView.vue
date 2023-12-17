<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			token: "",
			username: "",
			errorInUsername: true
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.token = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}

function usernameValidate(){
	alert("insert")
	return false
	this.errorUsername = this.username.length >= 3 &&
	this.username.length <= 16 &&
	/^.*?$/.test(this.username)
	
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
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						Sig-in/Log-in
					</button>
				</div>
			</div>
		</div>

		<div></div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="usernameValid" :msg="err"></ErrorMsg>
		
		<h5>Username</h5>
		<input v-model="username" placeholder="write hear" size="16" @input="usernameValidate">



		
	</div>
</template>

<style>
</style>

