import {createRouter, createWebHashHistory} from 'vue-router'
import Login from '../views/Login.vue'
import ProfileView from '../views/ProfileView.vue'
import ServerError from '../views/ServerError.vue'
import ClientError from '../views/ClientError.vue'
import { onBeforeUpdate } from 'vue'
const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/session'},
		{path: '/error/server', component: ServerError},
		{path: '/error/client', component: ClientError},
		{path: '/session', component: Login},
		{path: '/users/:id', component: ProfileView, meta:{
			requireAuth: true,
		}
	},
		
	]
})

export default router
