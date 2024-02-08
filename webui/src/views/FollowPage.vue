<script>
export default {
	data: function() {
		return {
			errormsg: null,
			positiveBanner: null,
			loading: false,
			some_data: null,
			userId: sessionStorage.getItem('userId'),
			username: sessionStorage.getItem('username'),
			logged: sessionStorage.getItem('logged')
		}
	},
	methods: {
		async followUser() {
            this.loading = true;
            this.errormsg = null;
			const toFollowUsername = document.getElementById("follow-username-input").value;
			document.getElementById("follow-username-input").value = ""
			console.log(this.username)
            try {
                let response = await this.$axios.post("/users/"+sessionStorage.getItem('userId')+"/followed/?username="+ toFollowUsername, {}, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
                console.log("User succesfully followed", response.data);
				this.positiveBanner = 'You are now following "' + toFollowUsername + '"'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

		async banUser() {
            this.loading = true;
            this.errormsg = null;
			const toBanUsername = document.getElementById("ban-username-input").value;
			document.getElementById("ban-username-input").value = ""
            try {
                let response = await this.$axios.post("/users/"+sessionStorage.getItem('userId')+"/banned/?username="+ toBanUsername, {}, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
                console.log("User succesfully followed", response.data);
				this.positiveBanner = '"' + toBanUsername + '" Has been banned'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

		async unbanUser() {
            this.loading = true;
            this.errormsg = null;
			const toUnbanUsername = document.getElementById("unban-username-input").value;
			document.getElementById("unban-username-input").value = ""
            try {
				console.log("/users/"+sessionStorage.getItem('userId')+"/banned/"+toUnbanUsername)
                let response = await this.$axios.delete("/users/"+sessionStorage.getItem('userId')+"/banned/"+toUnbanUsername, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
                console.log("User succesfully unbanned", response.data);
				this.positiveBanner = '"' + toUnbanUsername + '" Is not banned anymore'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
		async unfollowUser() {
            this.loading = true;
            this.errormsg = null;
			const toUnfollowUsername = document.getElementById("unfollow-username-input").value;
			document.getElementById("unfollow-username-input").value = ""
            try {
				console.log("/users/"+sessionStorage.getItem('userId')+"/followed/"+toUnfollowUsername)
                let response = await this.$axios.delete("/users/"+sessionStorage.getItem('userId')+"/followed/"+toUnfollowUsername, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
                console.log("User succesfully unfollowed", response.data);
				this.positiveBanner = 'You are not following "' + toUnfollowUsername + '" anymore'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }
	},
	mounted() {
		//this.refresh()
	}
}
</script>

<template>
	<p v-if="!this.logged" class="positive-banner" style="margin-top: 40px;">In this page you will be able to manage your network, but please log-in first in the home page</p>
	<div v-if="this.logged">
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Your Network</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<p v-if="positiveBanner" class="positive-banner">{{ this.positiveBanner }}</p>
		<p class="h3">Want to make new friends?</p>
		<div style="border: solid; border-color: grey; border-radius: 10px; padding: 20px; margin-bottom: 20px; margin-right: 40vw; background-color: rgba(208, 219, 252,1);" >
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="follow-username-input" style="margin-right: 10px;" class="form-control" placeholder="Who do you want to follow?">
        		<button class="btn btn-outline-secondary" type="button" @click="followUser">Follow</button>
			</div>
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="unfollow-username-input" style="margin-top: 20px; margin-right: 10px;" class="form-control" placeholder="Who do you want to unfollow?">
        		<button class="btn btn-outline-secondary" type="button" @click="unfollowUser" style="margin-top: 20px;">Unfollow</button>
			</div>
		</div>
		<p class="h3">Someone is bothering you?</p>
		<div style="border: solid; border-color: grey; border-radius: 10px; padding: 20px; margin-right: 40vw; background-color: rgba(100, 210, 180, 0.6);">
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="ban-username-input" style="margin-right: 10px;" class="form-control" placeholder="Who do you want to ban?">
        		<button class="btn btn-outline-secondary" type="button" @click="banUser">Ban</button>
			</div>	
				<div style="display: flex; flex-direction: row;">
				<input type="text" id="unban-username-input" style="margin-top: 20px; margin-right: 10px;" class="form-control" placeholder="Who do you want to unban?">
        	<button class="btn btn-outline-secondary" type="button" style=" margin-top: 20px;" @click="unbanUser">Unban</button>
			</div>
		</div>
	</div>
</template>

<style>
</style>

