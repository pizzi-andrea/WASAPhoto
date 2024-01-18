<script>
import { RouterLink, RouterView } from 'vue-router'
export default {
	props: ['usr'],

    data: function() {
        return {
            errormsg: "",
            id: this.usr.uid,
            username: this.usr.username,
            l_followers: [],
            followers: 0,
            isBanned: false,
            isFollower: false,
            followAllow: false,
        }
    },

    methods: {
        async putBan(){
            
            if (this.isBanned) {
                return
            }

            try{
                let response = {}
                response = await this.$axios.put("/users/" + localStorage.getItem('token')  + "/banned/" + this.id);
                switch(response.status){
                    case 201:
                        this.isBanned = true;
                        this.refresh();
                        break;
                    case 404:
                    
                    case 400:
                    
                    case 500:
                    
                    default:
                        console.log(response)
                }
            }catch(e){
                

            }


        },
        async refresh(){
            this.isBannedUser();
            this.getFollowers();

        },
        async putFollow(){
            if (this.isFollower){
                return
            }
            try{
                let response = {}
                response = await this.$axios.put("/users/" +  this.id + "/followers/" + localStorage.getItem('token'));
                switch(response.status){
                    case 201:
                        this.isFollower = true;
                        this.refresh();
                        break;
                        
                    case 404:
                    
                    case 400:
                    
                    case 500:
                    
                    default:
                        console.log(response)
                }
            }catch(e){
                

            }

            this.getFollowers()

        },

        async deleteFollow(){
            if( !this.isFollower){
                return
            }
            try{
                let response = {}
                response = await this.$axios.delete("/users/" +  this.id + "/followers/" + localStorage.getItem('token'));
                switch(response.status){
                    case 204:
                        this.isFollower =false;
                        this.refresh();
                        break;
                        case 404:
                    
                    case 400:
                    
                    case 500:
                    
                    default:
                        console.log(response)
                }
            }catch(e){
                

            }

        },

        async deleteBan(){
            if (!this.isBanned) {
                return
            }

            try{
                let response = {}
                response = await this.$axios.delete("/users/" + localStorage.getItem('token')  + "/banned/" + this.id);
                switch(response.status){
                    case 204:
                        this.isBanned = false;
                        this.refresh();
                        break;
                    case 404:
                    
                    case 400:
                    
                    case 500:
                    
                    default:
                        console.log(response)
                }
            }catch(e){
                

            }

        },

        async isBannedUser(){
            try{
                let response = {}
                response = await this.$axios.get("/users/" + localStorage.getItem('token')  + "/banned/" + this.id);
                switch(response.status){
                    case 200:
                        this.isBanned = true;
                        this.followAllow = true;
                        break;
                    
                    case 204:
                        this.isBanned = false;
                        this.followAllow = false;
                        break;
                    case 404:
                    
                    case 400:
                    
                    case 500:
                    
                    default:
                        console.log(response)
                }
            }catch(e){
                

            }

        },

        async getFollowers(){
            try{
                let response = {}
                response = await this.$axios.get("/users/" +  this.id + "/followers/");
                switch(response.status){
                    case 200:
                        this.l_followers = response.data;
                        this.followers =  this.l_followers.length
                        break;
                    
                    case 204:
                        this.isFollower = false;
                        break;
                    case 404:
                    
                    case 403:
                        this.followers = "?";
                        break;
                    
                    case 400:
                    
                    case 500:
                    
                    default:
                        console.log(response)
                }
            }catch(e){
                

            }

            this.isFollowerUser();

        },

        isFollowerUser(){
            let t= localStorage.getItem('token');

            for (const follower of this.l_followers){
                if(t == follower.uid){
                    this.isFollower = true;
                    return;
                }

            }
            this.isFollower= false;

        }




    },
    mounted() {
        this.refresh();

    },
    
    
}
</script>

<template>
	<li class="list-group-item d-flex justify-content-between align-items-start rounded m-2 ">
        <div class="ms-2 me-auto">
        
        <RouterLink :to= " '/users/' + this.id + '/profile'">
            <div class="fw-bold">{{username}}</div>  
        </RouterLink>
                               
        </div>
        <div class="btn-group ms-1 me-auto" role="group">
            <button type="button" class="btn btn-danger" style="--bs-btn-padding-y: .25rem; --bs-btn-padding-x: .5rem; --bs-btn-font-size: .75rem;" v-if="isFollower" @click="deleteFollow">
                Non seguire
                <span class="badge text-bg-primary rounded-pill">{{followers}}</span>
            </button>
            <button type="button" class="btn btn-primary" style="--bs-btn-padding-y: .25rem; --bs-btn-padding-x: .5rem; --bs-btn-font-size: .75rem;" v-else @click="putFollow" :disabled="followAllow" >
                Segui
                <span class="badge text-bg-primary rounded-pill">{{followers}}</span>
            </button>  
            <button type="button" class="btn btn-warning" style="--bs-btn-padding-y: .25rem; --bs-btn-padding-x: .5rem; --bs-btn-font-size: .75rem;" v-if="isBanned" @click="deleteBan">
                Sblocca
            </button>
            <button type="button" class="btn btn-danger" style="--bs-btn-padding-y: .25rem; --bs-btn-padding-x: .5rem; --bs-btn-font-size: .75rem;" v-else @click="putBan">
                Blocca
            </button>  
            </div>        
            
                                
    </li>
</template>

<style>
</style>



