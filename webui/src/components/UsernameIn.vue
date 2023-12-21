<script>
import { ref } from 'vue';

export default {
	data: function() {
		return {
			isValid: ref(false),
			msgError: ref("\
			Username format must have\
			 minimum 3 characters and maximum 16 characters"
			),
			username: ref("")
		}
	},

	methods: {
		async refresch(){
			this.isValid = !(this.username.length >= 3 &&
		   this.username.length <= 16 &&
		   /^.*$/.test(this.username) )
		}
	}
}

</script>

<template>
	<div>
		<input type="text" maxlength="16" minlength="1" pattern="^.*$" v-model="username" @input="refresch">
		<ErrorMsg v-if="isValid" :msg="msgError" ></ErrorMsg>
	</div>
</template>

<style>
</style>