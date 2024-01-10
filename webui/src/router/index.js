import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/Login.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session', component: Login},
		{path: '/users/:id', component: ProfileView},
		
	]
})

export default router
export const token = {}
