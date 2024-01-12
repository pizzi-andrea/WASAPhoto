import {createRouter, createWebHashHistory} from 'vue-router'
import Login from '../views/Login.vue'
import ProfileView from '../views/ProfileView.vue'
import ServerError from '../views/ServerError.vue'
import ClientError from '../views/ClientError.vue'
const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/error/server', component: ServerError},
		{path: '/error/client', component: ClientError},
		{path: '/session', component: Login},
		{path: '/users/:id', component: ProfileView},
		
	]
})

export default router
