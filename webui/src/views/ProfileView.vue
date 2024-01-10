<script>
import token from "./Login.vue"
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			User: null,
		}
	},
	methods: {
		async refresh() {
			try {
                let response = await this.$axios.get(this.$route.path, {
					headers: {
    					'Authorization': `Bearer ${token.Value}` 
  					}
				})
                this.User = response
            }
            catch(e) {
				this.errormsg = e.toString() + "Token: ";
				return 

            }
			
		},


	},
	mounted() {
		this.refresh()
	}
}



</script>

<template>
	<title>WasaPhoto</title>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Ciao, {{"ciao"}}</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="refresh">
						Aggiorna stream
					</button>
				</div>
			</div>
		</div>

		<div class="align-items-center">
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>		
	</div>
</template>

<style>
</style>

